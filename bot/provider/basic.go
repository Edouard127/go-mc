package provider

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/Edouard127/go-mc/bot/basic"
	"github.com/Edouard127/go-mc/bot/core"
	"github.com/Edouard127/go-mc/bot/screen"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/data/effects"
	"github.com/Edouard127/go-mc/data/enums"
	"github.com/Edouard127/go-mc/data/packetid"
	. "github.com/Edouard127/go-mc/data/slots"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/maths"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/Edouard127/go-mc/net/transactions"
	"github.com/google/uuid"
	"time"
)

type Player struct {
	world.PlayerInfo
	world.WorldInfo
	*core.EntityPlayer
	*core.Controller
	*screen.Manager
	*transactions.Transactions
	Settings             basic.Settings
	expBar               float32
	TotalExp             int32
	Level                int32
	isSpawn              bool
	fallTicks            float32
	fallDistance         float32
	stepHeight           float32
	jumpTicks            float32
	collidedHorizontally bool
	collidedVertically   bool
	collided             bool
}

func NewPlayer(settings basic.Settings) *Player {
	return &Player{
		PlayerInfo:   world.PlayerInfo{},
		WorldInfo:    world.WorldInfo{},
		EntityPlayer: core.NewEntityPlayer("", 0, uuid.UUID{}, 116, 0, 0, 0, 0, 0), // temporary
		Controller:   &core.Controller{},
		Manager:      screen.NewManager(),
		Transactions: transactions.NewTransactions(),
		Settings:     settings,
		isSpawn:      false,
	}
}

func (p *Player) GetExp() (float32, int32, int32) {
	return p.expBar, p.TotalExp, p.Level
}

func (p *Player) SetExp(bar float32, exp, total int32) {
	p.expBar, p.TotalExp, p.Level = bar, total, exp
}

func (p *Player) Respawn(c *Client) (err error) {
	const PerformRespawn = 0

	err = c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientCommand,
		pk.VarInt(PerformRespawn),
	))

	return
}

func (p *Player) Chat(c *Client, msg string) error {
	var salt int64
	binary.Read(rand.Reader, binary.BigEndian, &salt)

	var (
		message   = pk.String(msg[:min(len(msg), 256)])
		timestamp = pk.Long(time.Now().Unix())
	)

	err := c.Conn.WritePacket(pk.Marshal(packetid.SPacketChatMessage, message, timestamp, pk.Long(salt), pk.ByteArray(c.Auth.KeyPair.PublicKeySignatureV2), pk.Boolean(true)))
	if err != nil {
		return err
	}

	return nil
}

func RunTransactions(c *Client) error {
	if t := c.Player.Transactions.Next(); t == nil {
		return nil
	} else {
		for _, v := range t.Packets {
			if err := c.Conn.WritePacket(*v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Player) handleJumpWater() {
	p.Motion.Y += 0.03999999910593033
}

func (p *Player) handleJumpLava() {
	p.Motion.Y += 0.03999999910593033
}

func (p *Player) Jump() {
	p.Motion.Y += 0.42

	if p.IsPotionActive(effects.JumpBoost) {
		p.Motion.Y += 0.1 * float64(p.GetPotionEffect(effects.JumpBoost).Amplifier+1)
	}

	// TODO: Check if the player is sprinting
}

func (p *Player) travel(c *Client) error {
	result, err := c.World.GetBlock(p.Position)
	if err != nil {
		return err
	}

	switch result {
	case block.Water:
		//d0 := p.Position.Y
		f1 := enums.WaterInertia
		f2 := 0.02

		// TODO: Get depth strider enchantment level

		p.moveRelative(p.Motion, f2)
		p.move(enums.MoverTypeSelf, c, p.Motion)
		p.Motion.X *= f1
		p.Motion.Y *= 0.800000011920929
		p.Motion.Z *= f1

		p.Motion.Y -= 0.02

		if p.collidedHorizontally /*&& p.isOffsetPositionInLiquid(p.Motion.X, p.Motion.Y+0.6000000238418579-p.Position.Y+d0, p.Motion.Z)*/ {
			p.Motion.Y = 0.3
		}

	case block.Lava:
		//d4 := p.Position.Y
		p.moveRelative(p.Motion, 0.02)
		p.move(enums.MoverTypeSelf, c, p.Motion)

		p.Motion.X *= 0.5
		p.Motion.Y *= 0.5
		p.Motion.Z *= 0.5

		p.Motion.Y -= 0.02

		if p.collidedHorizontally /*&& p.isOffsetPositionInLiquid(p.Motion.X, p.Motion.Y+0.6000000238418579-p.Position.Y+d4, p.Motion.Z)*/ {
			p.Motion.Y = 0.30000001192092896
		}

	default:
		f6 := 0.91

		if p.OnGround {
			result, err = c.World.GetBlock(p.Position.Sub(maths.Vec3d{Y: 1}))
			if err != nil {
				return err
			}
			f6 = enums.Slipperiness(result.State())
		}

		//f7 := 0.16277136 / (f6 * f6 * f6)
		f8 := 0.02
		p.moveRelative(p.Motion, f8)
		f6 = 0.91 // FIXME: What is that

		// TODO: Check if the player is on a ladder

		p.move(enums.MoverTypeSelf, c, p.Motion)

		// TODO: Check if the player is on a ladder and collided horizontally

		if p.IsPotionActive(effects.Levitation) {
			p.Motion.Y += (0.05*float64(p.GetPotionEffect(effects.Levitation).Amplifier+1) - p.Motion.Y) * 0.2
		} else {
			p.Motion.Y -= 0.08
		}

		p.Motion.Y *= 0.98
		p.Motion.X *= f6
		p.Motion.Z *= f6
	}

	return nil
}

func (p *Player) move(moveType enums.MoverType, c *Client, motion maths.Vec3d) error {
	if moveType == enums.MoverTypePiston {
	}

	x := motion.X
	y := motion.Y
	z := motion.Z

	result, err := c.World.GetBlock(p.Position)
	if err != nil {
		return err
	}

	if result == block.Cobweb {
		x *= 0.25
		y *= 0.05000000074505806
		z *= 0.25
	}

	if (moveType == enums.MoverTypePlayer || moveType == enums.MoverTypeSelf) && p.OnGround {
		/*for ; x != 0.0 && len(c.World.GetCollisionBoxes(*p.UnaliveEntity, p.BoundingBox.Offset(x, -1.0, 0.0))) == 0; x = motion.X {
			if x < 0.05 && x >= -0.05 {
				x = 0.0
			} else if x > 0.0 {
				x -= 0.05
			} else {
				x += 0.05
			}
		}

		for ; z != 0.0 && len(c.World.GetCollisionBoxes(*p.UnaliveEntity, p.BoundingBox.Offset(0.0, -1.0, z))) == 0; y = motion.Y {
			if z < 0.05 && z >= -0.05 {
				z = 0.0
			} else if z > 0.0 {
				z -= 0.05
			} else {
				z += 0.05
			}
		}

		for ; x != 0.0 && z != 0.0 && len(c.World.GetCollisionBoxes(*p.UnaliveEntity, p.BoundingBox.Offset(x, -1.0, z))) == 0; z = motion.Z {
			if x < 0.05 && x >= -0.05 {
				x = 0.0
			} else if x > 0.0 {
				x -= 0.05
			} else {
				x += 0.05
			}

			x = motion.X

			if z < 0.05 && z >= -0.05 {
				z = 0.0
			} else if z > 0.0 {
				z -= 0.05
			} else {
				z += 0.05
			}
		}*/

		// TODO: Add entity collision

		//flag := p.OnGround || y != motion.Y && y < 0.0

		// TODO: Add step

		p.collidedHorizontally = x != motion.X || z != motion.Z
		p.collidedVertically = y != motion.Y
		p.OnGround = p.collidedVertically && y < 0.0
		p.collided = p.collidedHorizontally || p.collidedVertically
		//j6 := math.Floor(p.Position.X)
		//i1 := math.Floor(p.Position.Y - 0.20000000298023224)
		//k6 := math.Floor(p.Position.Z)
		//getBlock, _ := c.World.GetBlock(maths.Vec3d[float64]{X: j6, Y: i1, Z: k6})
		if p.OnGround {
			p.fallDistance = 0.0
		} else if y < 0.0 {
			p.fallDistance -= float32(y)
		}

		if x != motion.X {
			p.Motion.X = 0.0
		}

		if y != motion.Y {
			p.Motion.Y = 0.0
		}

		if z != motion.Z {
			p.Motion.Z = 0.0
		}
	}
	return nil
}

func (p *Player) moveRelative(motion maths.Vec3d, friction float64) {
}

func (p *Player) updateFallState(y float64, onGround bool, getBlock block.Block) {
	if onGround {
		p.fallDistance = 0.0
	} else if y < 0.0 {
		p.fallDistance -= float32(y)
	}
}

func ApplyPhysics(c *Client) error {
	if err := c.Conn.WritePacket(
		pk.Marshal(
			packetid.SPacketPlayerPosition,
			pk.Double(c.Player.Position.X),
			pk.Double(c.Player.Position.Y),
			pk.Double(c.Player.Position.Z),
			pk.Boolean(c.Player.OnGround),
		),
	); err != nil {
		return fmt.Errorf("failed to send player position: %v", err)
	}

	if c.Player.jumpTicks > 0 {
		c.Player.jumpTicks--
	}

	/*if c.Player.Motion.Length() < enums.NegligibleVelocity {
		c.Player.Motion = maths.NullVec3d
		return basic.Error{Err: basic.NoError, Info: nil}
	}*/

	if c.Player.Controller.Jump {
		result, err := c.World.GetBlock(c.Player.Position)
		if err != nil {
			return err
		}

		if result == block.Water || result == block.Lava {
			c.Player.Motion.Y += 0.03999999910593033
		} else {
			if c.Player.OnGround && c.Player.jumpTicks == 0 {
				c.Player.Jump()
				c.Player.jumpTicks = 10
			}
		}
	} else {
		c.Player.jumpTicks = 0
	}

	// Start section: Travel
	c.Player.travel(c)

	c.Player.Position = c.Player.Position.Add(c.Player.Motion)

	return nil
}

func (p *Player) WalkTo(c *Client, pos maths.Vec3d) error {
	path := c.World.PathFind(p.Position, pos)
	for _, v := range path {
		// Set the motion
		fmt.Println(v.Sub(p.Position))
		p.Motion = v.Sub(p.Position)
	}

	return nil
}

func (p *Player) ContainerClick(c *Client, id int, slot int16, button byte, mode int32, slots ChangedSlots, carried *Slot) error {
	return c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClickWindow,
		pk.UnsignedByte(id),
		pk.VarInt(p.Manager.StateID),
		pk.Short(slot),
		pk.Byte(button),
		pk.VarInt(mode),
		&slots,
		carried,
	))
}
