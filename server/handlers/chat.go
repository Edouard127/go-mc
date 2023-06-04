package handlers

import (
	"context"
	"fmt"
	"github.com/Tnze/go-mc/server/command"
)

var builder = command.NewGraph()

func init() {
	builder.AppendLiteral(builder.Literal("ban").
		AppendArgument(builder.Argument("target", command.StringParser(1)).
			HandleFunc(func(ctx context.Context, args []command.ParsedData) error {
				fmt.Println("The ban hammer has spoken!")
				return nil
			})).
		Unhandle(),
	).AppendLiteral(builder.Literal("ban-ip").
		AppendArgument(builder.Argument("target", command.StringParser(1)).
			HandleFunc(func(ctx context.Context, args []command.ParsedData) error {
				fmt.Println("The ban hammer has spoken!")
				return nil
			})).
		Unhandle(),
	)
}
