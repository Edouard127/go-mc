package provider

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"github.com/Edouard127/go-mc/auth/data"
	"github.com/Edouard127/go-mc/bot/basic"
	"github.com/Edouard127/go-mc/bot/core"
	"github.com/Edouard127/go-mc/bot/screen"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/data/effects"
	"github.com/Edouard127/go-mc/data/entity"
	"github.com/Edouard127/go-mc/data/enums"
	"github.com/Edouard127/go-mc/data/packetid"
	"github.com/Edouard127/go-mc/internal/util"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/maths"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/Edouard127/go-mc/net/transactions"
	"github.com/google/uuid"
	"time"
)

type Player struct {
	Settings     basic.Settings
	World        *world.World
	Controller   *core.Controller
	Manager      *screen.Manager
	PlayerInfo   world.PlayerInfo
	WorldInfo    world.WorldInfo
	EntityPlayer *core.EntityPlayer
	Transactions *util.Queue[[]*transactions.SlotAction]
	Abilities    *Abilities

	// Player info
	ExpBar               float64
	TotalExp             int32
	Level                int32
	IsSpawn              bool
	JumpTicks            int32
	JumpTriggerTime      int32
	JumpRidingTicks      int32
	JumpRidingScale      float64
	SprintTriggerTime    int32
	FallTicks            float64
	FallDistance         float64
	StepHeight           float64
	CollidedHorizontally bool
	CollidedVertically   bool
	Collided             bool
	UsingItem            bool
}

// NewPlayer creates a new Player
// The player should not have access to the Client struct
// but the player needs access to the world and the current
// authentication information
//
// The player should not be created before the client is
// created.
func NewPlayer(settings basic.Settings, clientWorld *world.World, info data.Auth) *Player {
	return &Player{
		Settings:     settings,
		World:        clientWorld,
		PlayerInfo:   world.PlayerInfo{},
		WorldInfo:    world.WorldInfo{},
		EntityPlayer: core.NewEntityPlayer(info.Name, 0, uuid.MustParse(info.UUID), 116, 0, 0, 0, 0, 0),
		Controller:   &core.Controller{},
		Manager:      screen.NewManager(),
		Transactions: util.NewQueue[[]*transactions.SlotAction](),
		Abilities:    NewAbilities(),
		IsSpawn:      false,
	}
}

func (pl *Player) Respawn(c *Client) (err error) {
	const PerformRespawn = 0

	err = c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientCommand,
		pk.VarInt(PerformRespawn),
	))

	return
}

func (pl *Player) UseItem(c *Client, hand enums.Hand) (err error) {
	return c.Conn.WritePacket(pk.Marshal(packetid.SPacketUseItem, pk.VarInt(hand)))
}

func (pl *Player) Chat(c *Client, msg string) error {
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

func RunTransactions(c *Client, cancel context.CancelFunc) error {
	var next []*transactions.SlotAction
	for c.Player.Transactions.Len() > 0 {
		next = c.Player.Transactions.Pop()

		for i := range next {
			c.Conn.WritePacket(pk.Marshal(packetid.SPacketClickWindow, pk.UnsignedByte(c.Player.Manager.CurrentScreen.GetType()), pk.VarInt(c.Player.Manager.StateID), next[i]))
		}
	}

	return nil
}

// Step is called every tick
// I don't have to code all the physic by myself
// thx minecrossoft 🙏
func Step(cl *Client, cancel context.CancelFunc) error {
	feetBlock, err := cl.Player.World.GetBlock(cl.Player.EntityPlayer.Position)
	if err != nil {
		return err
	}
	isJumping := cl.Player.Controller.Jump
	isSneaking := cl.Player.Controller.Sneak

	if cl.Player.UsingItem {
		cl.Player.Controller.LeftImpulse *= 0.2
		cl.Player.Controller.RightImpulse *= 0.2
	}

	if isSneaking {
		cl.Player.SprintTriggerTime = 0
	}

	flag := cl.Player.EntityPlayer.Food > 6 && cl.Player.Abilities.AllowFlying
	if !cl.Player.Controller.Sprint && !cl.Player.UsingItem && !cl.Player.EntityPlayer.IsPotionActive(effects.Blindness) {
		if flag && !cl.Player.Controller.Sneak {
			cl.Player.SprintTriggerTime = 7
		} else {
			cl.Player.SprintTriggerTime = 0
		}
	}

	if feetBlock == block.Water && cl.Player.Controller.Sneak && !cl.Player.Abilities.Flying {
		cl.Player.EntityPlayer.Motion = cl.Player.EntityPlayer.Motion.AddScalar(0, -0.04, 0)
	}

	if cl.Player.Abilities.Flying {
		var mul float64
		if cl.Player.Controller.Sneak {
			mul = -1
		} else {
			mul = 1
		}

		if mul != 0 {
			cl.Player.EntityPlayer.Motion = cl.Player.EntityPlayer.Motion.AddScalar(0, mul*cl.Player.Abilities.FlyingSpeed*3.0, 0)
		}
	}

	if cl.Player.RidingJumpable() {
		riding := cl.Player.EntityPlayer.Vehicle
		if cl.Player.JumpRidingTicks < 0 {
			cl.Player.JumpRidingTicks++
			if cl.Player.JumpRidingTicks == 0 {
				cl.Player.JumpRidingScale = 0
			}
		}

		if isJumping && !cl.Player.Controller.Jump {
			cl.Player.JumpRidingTicks++
			if cl.Player.JumpRidingTicks < 10 {
				cl.Player.JumpRidingScale = float64(cl.Player.JumpRidingTicks) * 0.1
			} else {
				cl.Player.JumpRidingScale = 0.8 + 2.0/float64(cl.Player.JumpRidingTicks-9)*0.1
			}
			cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerCommand, pk.VarInt(riding.GetID()), pk.VarInt(enums.StartJumpWithHorse), pk.VarInt(cl.Player.JumpRidingScale*100)))
		} else if !isJumping && cl.Player.Controller.Jump {
			cl.Player.JumpRidingTicks = 0
			cl.Player.JumpRidingScale = 0
		} else if isJumping {
			cl.Player.JumpRidingTicks++
			if cl.Player.JumpRidingTicks < 10 {
				cl.Player.JumpRidingScale = float64(cl.Player.JumpRidingTicks) * 0.1
			} else {
				cl.Player.JumpRidingScale = 0.8 + 2.0/float64(cl.Player.JumpRidingTicks-9)*0.1
			}
		}
	} else {
		cl.Player.JumpRidingScale = 0
	}

	if cl.Player.EntityPlayer.OnGround && cl.Player.Abilities.Flying && cl.Player.PlayerInfo.Gamemode != uint8(enums.Spectator) {
		cl.Player.Abilities.Flying = false
		cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerAbilities, cl.Player.Abilities))
	}

	return nil
}

func (pl *Player) move(mover enums.MoverType, pos0 maths.Vec3d) {
	feet, err := pl.World.GetBlock(pl.EntityPlayer.Position)
	if err != nil {
		return
	}

	eyes, err := pl.World.GetBlock(pl.EntityPlayer.EyePosition)
	if err != nil {
		return
	}

	if mover == enums.MoverTypePiston {
		pos0 = limitPistonMovement(pos0)
	}

	if feet != block.Air && eyes != block.Air {
		pl.EntityPlayer.Motion = maths.Vec3d{}
	}

	pl.backOffFromEdge(pos0, mover)

	pl.CollidedHorizontally = pl.World.CollidesHorizontally(pl.EntityPlayer.BoundingBox)
	pl.CollidedVertically = pl.World.CollidesVertically(pl.EntityPlayer.BoundingBox)

	pl.EntityPlayer.OnGround = pl.CollidedVertically && pl.EntityPlayer.Position.Y < 0

	/*onPos, err := pl.World.GetBlock(pl.EntityPlayer.Position.SubScalar(0, 0.2, 0))
	if err != nil {
		return
	}*/

	if pl.CollidedHorizontally {
		pl.EntityPlayer.SetMotion(0, pl.EntityPlayer.Motion.Y, 0) // For now we will assume that the player collides with both x and z
	}

	// TODO: logic shit and whatever

	// Now let's do the fun stuff :D
	pl.EntityPlayer.Motion = pl.EntityPlayer.Motion.MulScalar(feet.SpeedFactor, 1, feet.SpeedFactor)
}

func (pl *Player) jumpPower() float64 {
	return 0.42*pl.World.MustGetBlock(pl.EntityPlayer.Position).JumpFactor + float64(pl.EntityPlayer.ActivePotionEffects[effects.JumpBoost.ID].Amplifier+1)
}

func (pl *Player) RidingJumpable() bool {
	return pl.EntityPlayer.IsPassenger() && pl.EntityPlayer.Vehicle.Equals(entity.Horse)
}

func limitPistonMovement(pos0 maths.Vec3d) maths.Vec3d {
	if pos0.LengthSquared() <= 1e-7 {
		return pos0
	} else {
		return pos0 // TODO
	}
}

func (pl *Player) backOffFromEdge(pos0 maths.Vec3d, mover enums.MoverType) maths.Vec3d {
	if pl.Abilities.Flying && pos0.Y <= 0 && (mover == enums.MoverTypeSelf || mover == enums.MoverTypePlayer) && pl.Controller.Sneak /*&& isAboveGround */ {
		return pos0
	}

	d0 := pos0.X
	d1 := pos0.Z

	for d0 != 0 && pl.World.CollidesWithAnyBlock(pl.EntityPlayer.GetBoundingBox().Move(d0, -1, 0)) {
		if d0 < 0.05 && d0 >= -0.05 {
			d0 = 0
		} else if d0 > 0 {
			d0 -= 0.05
		} else {
			d0 += 0.05
		}
	}

	for d1 != 0 && pl.World.CollidesWithAnyBlock(pl.EntityPlayer.GetBoundingBox().Move(0, -1, d1)) {
		if d1 < 0.05 && d1 >= -0.05 {
			d1 = 0
		} else if d1 > 0 {
			d1 -= 0.05
		} else {
			d1 += 0.05
		}
	}

	for d0 != 0 && d1 != 0 && pl.World.CollidesWithAnyBlock(pl.EntityPlayer.GetBoundingBox().Move(d0, -1, d1)) {
		if d0 < 0.05 && d0 >= -0.05 {
			d0 = 0
		} else if d0 > 0 {
			d0 -= 0.05
		} else {
			d0 += 0.05
		}

		if d1 < 0.05 && d1 >= -0.05 {
			d1 = 0
		} else if d1 > 0 {
			d1 -= 0.05
		} else {
			d1 += 0.05
		}
	}

	return maths.Vec3d{X: d0, Y: pos0.Y, Z: d1}
}
