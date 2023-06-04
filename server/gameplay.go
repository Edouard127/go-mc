package server

import (
	_ "embed"
	"github.com/Edouard127/go-mc/server/auth"
	"github.com/google/uuid"

	"github.com/Edouard127/go-mc/net"
)

type GamePlay interface {
	// AcceptPlayer handle everything after "LoginSuccess" is sent.
	//
	// Note: the connection will be closed after this function returned.
	// You don't need to close the connection, but to keep not returning while the player is playing.
	AcceptPlayer(name string, id uuid.UUID, profilePubKey *auth.PublicKey, properties []auth.Property, protocol int32, conn *net.Conn)
}
