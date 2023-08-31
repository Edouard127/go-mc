package provider

import (
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
)

type Abilities struct {
	Invulnerable bool
	Flying       bool
	AllowFlying  bool
	InstantBuild bool
	CanBuild     bool
	FlyingSpeed  float64
	WalkingSpeed float64
}

func NewAbilities() *Abilities {
	return &Abilities{CanBuild: true, FlyingSpeed: 0.05, WalkingSpeed: 0.1}
}

func (a *Abilities) ReadFrom(r io.Reader) (int64, error) {
	var flags pk.Byte
	var flyingSpeed pk.Float
	nn, err := pk.Tuple{flags, flyingSpeed}.ReadFrom(r)
	if err != nil {
		return nn, err
	}

	a.FlyingSpeed = float64(flyingSpeed)

	a.Invulnerable = flags&0x01 != 0
	a.Flying = flags&0x02 != 0
	a.AllowFlying = flags&0x04 != 0
	a.InstantBuild = flags&0x08 != 0

	return nn, nil
}

func (a *Abilities) WriteTo(w io.Writer) (int64, error) {
	var flags byte
	if a.Flying {
		flags |= 0x02
	}
	return pk.Byte(flags).WriteTo(w)
}
