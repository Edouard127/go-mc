package command

import (
	"io"
	"strings"
	"unsafe"

	pk "github.com/Edouard127/go-mc/net/packet"
)

const (
	isExecutable = 1 << (iota + 2)
	hasRedirect
	hasSuggestionsType
)

func (g *Graph) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Array(g.nodes),
		pk.VarInt(0),
	}.WriteTo(w)
}

func (n Node) WriteTo(w io.Writer) (int64, error) {
	var flag byte
	flag |= n.kind & 0x03
	if n.Run != nil {
		flag |= isExecutable
	}
	return pk.Tuple{
		pk.Byte(flag),
		pk.Array((*[]pk.VarInt)(unsafe.Pointer(&n.Children))),
		pk.Optional[pk.Boolean]{
			Has:   func() bool { return n.kind&hasRedirect != 0 },
			Value: nil, // TODO: send redirect node
		},
		pk.Optional[pk.Boolean]{
			Has:   func() bool { return n.kind == ArgumentNode || n.kind == LiteralNode },
			Value: pk.String(n.Name),
		},
		pk.Optional[pk.Boolean]{
			Has:   func() bool { return n.kind == ArgumentNode },
			Value: n.Parser, // Parser identifier and Properties
		},
		pk.Optional[pk.Boolean]{
			Has:   func() bool { return flag&hasSuggestionsType != 0 },
			Value: nil, // TODO: send Suggestions type
		},
	}.WriteTo(w)
}

type ChatCommand struct {
	Command   pk.String
	Timestamp pk.Long
	Salt      pk.Long
	Args      []CommandArgument
	Signed    pk.Boolean
}

type CommandArgument struct {
	Name      pk.String
	Signature pk.ByteArray
}

func (c ChatCommand) WriteTo(w io.Writer) (int64, error) {
	var args []pk.Tuple
	for _, v := range c.Args {
		args = append(args, pk.Tuple{
			v.Name,
			pk.VarInt(len(v.Signature)),
			v.Signature,
		})
	}
	return pk.Tuple{
		c.Command,
		c.Timestamp,
		c.Salt,
		pk.VarInt(len(args)),
		args,
		c.Signed,
	}.WriteTo(w)
}

func (c *ChatCommand) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		&c.Command,
		&c.Timestamp,
		&c.Salt,
		&c.Args,
		&c.Signed,
	}.ReadFrom(r)
}

func (c *ChatCommand) String() string {
	var s strings.Builder
	s.WriteString(string(c.Command))
	for _, v := range c.Args {
		s.WriteByte(' ')
		s.WriteString(string(v.Name))
	}
	return s.String()
}

func (a CommandArgument) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		a.Name,
		pk.VarInt(len(a.Signature)),
		a.Signature,
	}.WriteTo(w)
}

func (a *CommandArgument) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		&a.Name,
		&a.Signature,
	}.ReadFrom(r)
}
