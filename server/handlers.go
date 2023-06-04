package server

import (
	"context"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server/command"
)

func ChatCommand(s *Server, p pk.Packet) error {
	var (
		chatCommand command.ChatCommand
	)

	if err := p.Scan(&chatCommand); err != nil {
		return err
	}
	err := s.Commands.Execute(context.TODO(), chatCommand.String())
	return err
}
