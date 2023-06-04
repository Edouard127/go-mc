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
	"github.com/Edouard127/go-mc/chat"
	"github.com/Edouard127/go-mc/data/packetid"
	"github.com/Edouard127/go-mc/net"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/Edouard127/go-mc/server/auth"
	"github.com/Edouard127/go-mc/server/client"
	"github.com/Edouard127/go-mc/server/command"
	"github.com/Edouard127/go-mc/server/world"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"path/filepath"
)

const ProtocolName = "1.19"
const ProtocolVersion = 759

type Server struct {
	*zap.Logger
	ListPingHandler
	LoginHandler
	World          *world.World
	PlayerList     *PlayerList
	Keepalive      *KeepAlive
	Commands       *command.Graph
	playerProvider *world.PlayerProvider
	settings       ServerSettings
}

func NewServer(settings ServerSettings) *Server {
	logger, _ := zap.NewDevelopment()
	return &Server{
		Logger:       logger.With(zap.String("module", "server")),
		LoginHandler: NewMojangLoginHandler(),
		World: world.NewWorld(
			logger.With(zap.String("module", "world")),
			"minecraft:overworld", // TODO: make this configurable
			world.NewProvider(filepath.Join(filepath.Join(".", settings.LevelName), "region"), settings.ChunkLoadingLimiter.Limiter()),
			world.Config{},
		),
		PlayerList:     NewPlayerList(settings.MaxPlayers),
		Keepalive:      NewKeepAlive(),
		Commands:       command.NewGraph(),
		playerProvider: world.NewPlayerProvider(filepath.Join(".", settings.LevelName, "playerdata")),
		settings:       settings,
	}
}

// Listen starts listening on the specified address.
func (s *Server) Listen(addr string) error {
	if addr == "" {
		addr = s.settings.ListenAddress
	}
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
			s.Logger.Error("error closing connection", zap.Error(err))
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
			s.Error("error accepting login", zap.Error(err))
			return
		}
		s.AcceptPlayer(name, id, profilePubKey, properties, protocol, conn, PlayerSample{Name: name, ID: id})
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

func (s *Server) Description() string {
	return s.settings.MessageOfTheDay
}

func (s *Server) FavIcon() string {
	return s.settings.Icon
}

func (s *Server) BroadcastNewPlayer(c *client.ServerClient, sample PlayerSample) {
	s.PlayerList.ClientJoin(c, sample)
	s.Keepalive.ClientJoin(c)
}

func (s *Server) AcceptPlayer(name string, id uuid.UUID, profilePubKey *auth.PublicKey, properties []auth.Property, protocol int32, conn *net.Conn, sample PlayerSample) {
	logger := s.With(
		zap.String("player", name),
		zap.String("uuid", id.String()),
	)
	c := client.NewServerClient(
		logger,
		conn, world.NewPlayer(name, id, profilePubKey, properties))

	p, err := s.playerProvider.GetPlayer(name, id, profilePubKey, properties)
	if err != nil {
		s.Error("error getting player", zap.Error(err))
		return
	}

	logger.Info("Player joined", zap.Int32("entity id", p.EntityID))
	defer logger.Info("Player left")

	s.BroadcastNewPlayer(c, sample)
	c.SendLogin(s.World)
	c.SendServerData(chat.TextPtr(s.Description()), s.FavIcon(), s.settings.EnforceSecureProfile)

	/* ... */
}
