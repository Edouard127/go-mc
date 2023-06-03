// Package server provide a minecraft server framework.
// You can build the server you want by combining the various functional modules provided here.
// An example can be found in examples/frameworkServer.
//
// # This package is under rapid development, and any API may be subject to break changes
//
// A server is roughly divided into two parts: Gate and GamePlay
//
//	+---------------------------------------------------------------------+
//	|                        Go-MC Server Framework                       |
//	+--------------------------------------+------------------------------+
//	|               Gate                   |           GamePlay           |
//	+--------------------+-----------------+                              |
//	|    LoginHandler    | ListPingHandler |                              |
//	+--------------------+------------+----+---------------+--------------+
//	| MojangLoginHandler |  PingInfo  |     PlayerList     |  Others....  |
//	+--------------------+------------+--------------------+--------------+
//
// Gate, which is used to respond to the client login request, provide login verification,
// respond to the List Ping Request and providing the online players' information.
//
// Gameplay, which is used to handle all things after a player successfully logs in
// (that is, after the LoginSuccess package is sent),
// and is responsible for functions including player status, chunk management, keep alive, chat, etc.
//
// The implement of Gameplay is provided at https://github.com/go-mc/server.
package server

import (
	"errors"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"log"
)

const ProtocolName = "1.19"
const ProtocolVersion = 759

type Server struct {
	*log.Logger
	ListPingHandler
	LoginHandler
	GamePlay
	Queue      *PacketQueue
	PlayerList *PlayerList
	Keepalive  *KeepAlive
	settings   ServerSettings
}

func NewServer(settings ServerSettings) *Server {
	return &Server{
		Logger:       log.Default(),
		LoginHandler: NewMojangLoginHandler(),
		Queue:        NewPacketQueue(),
		PlayerList:   NewPlayerList(settings.MaxPlayers),
		Keepalive:    NewKeepAlive(),
		settings:     settings,
	}
}

// Listen starts listening on the specified address.
func (s *Server) Listen(addr string) error {
	listener, err := net.ListenMC(addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.acceptConn(&conn)
	}
}

func (s *Server) acceptConn(conn *net.Conn) {
	defer func(conn *net.Conn) {
		err := conn.Close()
		if err != nil {
			s.Println(err)
		}
	}(conn)
	protocol, intention, err := s.handshake(conn)
	if err != nil {
		return
	}

	switch intention {
	case 1: // list ping
		s.acceptListPing(conn)
	case 2: // login
		name, id, profilePubKey, properties, err := s.AcceptLogin(conn, protocol)
		if err != nil {
			var loginErr *LoginFailErr
			if errors.As(err, &loginErr) {
				if err := conn.WritePacket(pk.Marshal(
					packetid.CPacketDisconnect,
					loginErr.reason,
				)); err != nil {
					return
				}
			}
			if s.Logger != nil {
				s.Logger.Printf("client %v login error: %v", conn.Socket.RemoteAddr(), err)
			}
			return
		}
		s.AcceptPlayer(name, id, profilePubKey, properties, protocol, conn)
	}
}

func (s *Server) Name() string {
	return s.settings.Name
}

func (s *Server) Protocol() int {
	return ProtocolVersion
}

func (s *Server) MaxPlayer() int {
	return s.settings.MaxPlayers
}

func (s *Server) OnlinePlayer() int {
	return s.PlayerList.OnlinePlayer()
}

func (s *Server) PlayerSamples() []PlayerSample {
	return s.PlayerList.PlayerSamples()
}

func (s *Server) Description() *chat.Message {
	return s.settings.MOTD
}

func (s *Server) FavIcon() string {
	return s.settings.Icon
}
