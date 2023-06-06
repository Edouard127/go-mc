package core

import (
	"github.com/Edouard127/go-mc/chat"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/google/uuid"
	"io"
)

type PlayerList struct {
	Players map[uuid.UUID]PlayerEntry
}

func NewPlayerList() *PlayerList {
	return &PlayerList{
		Players: make(map[uuid.UUID]PlayerEntry),
	}
}

func (p PlayerList) ReadFrom(r io.Reader) (int64, error) {
	var action pk.VarInt
	var uuids []pk.UUID
	var players []PlayerEntry

	n, err := action.ReadFrom(r)
	if err != nil {
		return n, err
	}

	var n1 int64

	if action != 0 {
		n1, err = pk.Array(&uuids).ReadFrom(r)
		if err != nil {
			return n + n1, err
		}
	} else {
		n1, err = pk.Array(&players).ReadFrom(r)
		if err != nil {
			return n + n1, err
		}
	}

	switch action {
	case 0:
		p.AddPlayers(players)
	case 1:
		gamemodes := make([]pk.VarInt, len(uuids))
		for i := range gamemodes {
			n2, _ := gamemodes[i].ReadFrom(r)
			n1 += n2
			p.UpdateGamemode(uuid.UUID(uuids[i]), int32(gamemodes[i]))
		}
	case 2:
		ping := make([]pk.VarInt, len(uuids))
		for i := range ping {
			n2, _ := ping[i].ReadFrom(r)
			n1 += n2
			p.UpdatePing(uuid.UUID(uuids[i]), int32(ping[i]))
		}
	case 3:
		displayName := make([]pk.Option[chat.Message, *chat.Message], len(uuids))
		for i := range displayName {
			n2, _ := displayName[i].ReadFrom(r)
			n1 += n2
			p.UpdateDisplayName(uuid.UUID(uuids[i]), displayName[i])
		}
	case 4:
		p.RemovePlayers(uuids)
	}

	return n + n1, nil
}

func (p *PlayerList) AddPlayers(players []PlayerEntry) {
	for _, v := range players {
		p.Players[v.UUID] = v
	}
}

func (p *PlayerList) RemovePlayers(uuids []pk.UUID) {
	for _, v := range uuids {
		delete(p.Players, uuid.UUID(v))
	}
}

func (p *PlayerList) GetPlayer(uuid uuid.UUID) PlayerEntry {
	return p.Players[uuid]
}

func (p *PlayerList) UpdateGamemode(uuid uuid.UUID, gamemode int32) {
	if _, ok := p.Players[uuid]; !ok {
		return
	}
	entry := p.Players[uuid]
	entry.Gamemode = gamemode
	p.Players[uuid] = entry
}

func (p *PlayerList) UpdatePing(uuid uuid.UUID, ping int32) {
	if _, ok := p.Players[uuid]; !ok {
		return
	}
	entry := p.Players[uuid]
	entry.Ping = ping
	p.Players[uuid] = entry
}

func (p *PlayerList) UpdateDisplayName(uuid uuid.UUID, displayName pk.Option[chat.Message, *chat.Message]) {
	if _, ok := p.Players[uuid]; !ok {
		return
	}
	entry := p.Players[uuid]
	entry.DisplayName = displayName
	p.Players[uuid] = entry
}

type PlayerEntry struct {
	UUID        uuid.UUID
	Name        string
	Properties  []pk.Property
	Gamemode    int32
	Ping        int32
	DisplayName pk.Option[chat.Message, *chat.Message]
	//Timestamp   int64
	/*PublicKey   []byte
	Signature   []byte*/
}

func (p *PlayerEntry) ReadFrom(r io.Reader) (int64, error) {
	n, err := pk.Tuple{
		(*pk.UUID)(&p.UUID),
		(*pk.String)(&p.Name),
		pk.Array(&p.Properties),
		(*pk.VarInt)(&p.Gamemode),
		(*pk.VarInt)(&p.Ping),
		&p.DisplayName,
	}.ReadFrom(r)
	return n, err
}
