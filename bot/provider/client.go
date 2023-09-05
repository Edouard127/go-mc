package provider

import (
	"github.com/Edouard127/go-mc/auth/data"
	"github.com/Edouard127/go-mc/bot/basic"
	"github.com/Edouard127/go-mc/bot/core"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/net"
)

// Client is used to access Minecraft server
//
//	+--------------------------------+-------------------------------+
//	|                        Client framework                        |
//	+--------------------------------+-------------------------------+
//	|       World       |       PlayerList       | 	    Player       |
//	+-------------------+------------+-------------------------------+
type Client struct {
	Conn *net.Conn
	Auth data.Auth

	World      *world.World
	PlayerList *core.PlayerList
	Player     *Player

	Events      Events[Client]
	LoginPlugin map[string]func(data []byte) ([]byte, error)
}

func (cl *Client) Close() error {
	return cl.Conn.Close()
}

// NewClient creates a new Client
// If you wish to use online-mode, refer to microsoft.LoginFromCache and microsoft.MinecraftLogin
func NewClient(auth data.Auth) *Client {
	clientWorld := world.NewWorld()
	return Attach(&Client{
		Auth:       auth,
		World:      clientWorld,
		PlayerList: core.NewPlayerList(),
		Player:     NewPlayer(basic.DefaultSettings, clientWorld, auth),
		Events:     Events[Client]{handlers: make(map[int32]*handlerHeap[Client]), tickers: new(tickerHeap[Client])},
	})
}
