package server

import "github.com/Tnze/go-mc/chat"

type ServerSettings struct {
	// The name of the server.
	Name string
	// The maximum number of players that can play on the server at the same time.
	MaxPlayers int
	// The server's MOTD.
	MOTD *chat.Message
	// The server's icon.
	Icon string
}

var DefaultServerSettings = ServerSettings{
	Name:       "Go-MC Server",
	MaxPlayers: 20,
	MOTD:       &chat.Message{Text: "Go-MC Server"},
	Icon:       "",
}
