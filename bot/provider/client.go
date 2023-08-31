package provider

import (
	"github.com/Edouard127/go-mc/auth/data"
	"github.com/Edouard127/go-mc/bot/basic"
	"github.com/Edouard127/go-mc/bot/core"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/net"
)

// Client is used to access Minecraft server
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
// By default, the authentication is offline-mode.
// If you wish to use online-mode, refer to microsoft.LoginFromCache and microsoft.MinecraftLogin
func NewClient() *Client {
	w := world.NewWorld()
	return Attach(&Client{
		Auth:       data.Auth{Profile: data.DefaultProfile}, // Offline-mode by default
		World:      w,
		PlayerList: core.NewPlayerList(),
		Player:     NewPlayer(basic.DefaultSettings, w),
		Events:     Events[Client]{handlers: make(map[int32]*handlerHeap[Client]), tickers: new(tickerHeap[Client])},
	})
}
