package client

import (
	"encoding/binary"
	"fmt"
	"github.com/Edouard127/go-mc/chat"
	"github.com/Edouard127/go-mc/data/packetid"
	"github.com/Edouard127/go-mc/net"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/Edouard127/go-mc/server/network"
	"github.com/Edouard127/go-mc/server/world"
	"go.uber.org/zap"
)

// ServerClient represents a client connected to the server.
type ServerClient struct {
	log      *zap.Logger
	conn     *net.Conn
	player   *world.Player
	queue    *network.PacketQueue
	handlers *network.Events[ServerClient]
}

func NewServerClient(log *zap.Logger, conn *net.Conn, player *world.Player) *ServerClient {
	return &ServerClient{
		log:      log,
		conn:     conn,
		player:   player,
		queue:    network.NewPacketQueue(), // TODO: Limit queue size
		handlers: network.NewEvents[ServerClient](),
	}
}

func (c *ServerClient) Start() {
	stopped := make(chan struct{}, 2)
	done := func() {
		stopped <- struct{}{}
	}
	// Exit when any error is thrown
	go c.startSend(done)
	go c.startReceive(done)
	<-stopped
}

func (c *ServerClient) startSend(done func()) {
	defer done()
	for {
		p, ok := c.queue.Pull()
		if !ok {
			return
		}
		err := c.conn.WritePacket(p)
		if err != nil {
			c.log.Debug("Send packet fail", zap.Error(err))
			return
		}
	}
}

func (c *ServerClient) startReceive(done func()) {
	defer done()
	var err error
	var p pk.Packet
	for {
		err = c.conn.ReadPacket(&p)
		if err != nil {
			c.log.Debug("Receive packet fail", zap.Error(err))
			return
		}
		c.log.Debug("Received packet", zap.String("id", fmt.Sprintf("%#x", p.ID)))
		err = c.handlePacket(p)
		if err != nil {
			c.log.Debug("Handle packet fail", zap.Error(err))
			return
		}
	}
}

func (c *ServerClient) handlePacket(p pk.Packet) error {
	for _, handler := range *c.handlers.GetHandlers(p.ID) {
		err := handler.F(c, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ServerClient) SendPacket(packet pk.Packet) {
	c.queue.Push(packet)
}

func (c *ServerClient) SendKeepAlive(id int64) {
	c.SendPacket(pk.Marshal(packetid.CPacketKeepAlive, pk.Long(id)))
}

func (c *ServerClient) SendDisconnect(reason chat.Message) {
	c.SendPacket(pk.Marshal(packetid.CPacketDisconnect, reason))
	c.conn.Close()
}

func (c *ServerClient) SendLogin(w *world.World) {
	hashedSeed := w.HashedSeed()
	c.SendPacket(pk.Marshal(packetid.CPacketLogin, pk.Int(c.player.EntityID), pk.String(c.player.Name), pk.Byte(-1), pk.Array([]pk.Identifier{
		pk.Identifier(w.Name),
	}), pk.NBT(world.NetworkCodec), pk.Identifier("minecraft:overworld"), pk.Identifier(w.Name), pk.Long(binary.BigEndian.Uint64(hashedSeed[:8])),
		pk.Boolean(true), pk.VarInt(0), // Max players (ignored by client)
		pk.VarInt(c.player.ViewDistance), // View Distance
		pk.VarInt(c.player.ViewDistance), // Simulation Distance
		pk.Boolean(false),                // Reduced Debug Info
		pk.Boolean(false),                // Enable respawn screen
		pk.Boolean(false),                // Is Debug
		pk.Boolean(false),                // Is Flat
		pk.Boolean(false),                // Has Last Death Location
	))
}

func (c *ServerClient) SendServerData(motd *chat.Message, favIcon string, enforceSecureProfile bool) {
	c.SendPacket(
		pk.Marshal(packetid.CPacketServerData, motd, pk.String(favIcon), pk.Boolean(enforceSecureProfile)),
	)
}
