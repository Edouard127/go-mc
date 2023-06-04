package server

import (
	"context"
	"fmt"
	"github.com/Tnze/go-mc/server/command"
)

func Attach(s *Server) {
	s.Commands.AppendLiteral(s.Commands.Literal("ban").
		AppendArgument(s.Commands.Argument("target", command.StringParser(1)).
			HandleFunc(func(ctx context.Context, args []command.ParsedData) error {
				fmt.Println("The ban hammer has spoken!")
				return nil
			})).
		Unhandle(),
	).AppendLiteral(s.Commands.Literal("ban-ip").
		AppendArgument(s.Commands.Argument("target", command.StringParser(1)).
			HandleFunc(func(ctx context.Context, args []command.ParsedData) error {
				fmt.Println("The ban hammer has spoken!")
				return nil
			})).
		Unhandle(),
	)
}
