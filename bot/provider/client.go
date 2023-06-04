package provider

import (
	"github.com/Edouard127/go-mc/bot/basic"
	"github.com/Edouard127/go-mc/bot/maths"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/net"
	auth "github.com/maxsupermanhd/go-mc-ms-auth"
)

// Client is used to access Minecraft server
type Client struct {
	Conn *net.Conn
	Auth Auth

	World  *world.World
	Player *Player
	TPS    *maths.TpsCalculator

	EventHandlers EventsListener
	Events        Events
	LoginPlugin   map[string]func(data []byte) ([]byte, error)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() *Client {
	c := &Client{
		Auth:   Auth{RSAuth: auth.RSAuth{Name: "Steve"}, KeyPair: auth.KeyPair{}},
		World:  world.NewWorld(),
		Player: NewPlayer(basic.DefaultSettings),
		TPS:    new(maths.TpsCalculator),
		Events: Events{handlers: make(map[int32]*handlerHeap)},
	}
	c.TPS.SetCallback(c.handleTickers)
	return c
}
