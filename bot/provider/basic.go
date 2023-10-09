package provider

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
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
	"math"
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
	ExpBar                    float64
	TotalExp                  int32
	Level                     int32
	IsSpawn                   bool
	JumpTicks                 int32
	JumpTriggerTime           int32
	JumpRidingTicks           int32
	JumpRidingScale           float64
	SprintTriggerTime         int32
	FallTicks                 float64
	FallDistance              float64
	StepHeight                float64
	PositionReminder          int
	CollidedHorizontally      bool
	CollidedHorizontallyMinor bool
	CollidedVertically        bool
	CollidedVerticallyBelow   bool
	Collided                  bool
	UsingItem                 bool
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
	isJumping := cl.Player.Controller.Jump
	isSneaking := cl.Player.Controller.Sneak

	if cl.Player.UsingItem {
		cl.Player.Controller.LeftImpulse *= 0.2
		cl.Player.Controller.ForwardImpulse *= 0.2
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

	if cl.Player.Abilities.Flying {
		var mul float64
		if cl.Player.Controller.Sneak {
			mul = -1
		} else {
			mul = 1
		}

		if mul != 0 {
			cl.Player.EntityPlayer.Motion.AddScalar(0, mul*cl.Player.Abilities.FlyingSpeed*3.0, 0)
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
		cl.Player.FallTicks = 0
		cl.Player.Abilities.Flying = false
		//cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerAbilities, cl.Player.Abilities))
	} else if !cl.Player.EntityPlayer.OnGround && cl.Player.PlayerInfo.Gamemode != uint8(enums.Spectator) {
		cl.Player.FallTicks++
	}

	if math.Abs(cl.Player.EntityPlayer.Motion.X) < 0.003 {
		cl.Player.EntityPlayer.Motion.X = 0
	}

	if math.Abs(cl.Player.EntityPlayer.Motion.Y) < 0.003 {
		cl.Player.EntityPlayer.Motion.Y = 0
	}

	if math.Abs(cl.Player.EntityPlayer.Motion.Z) < 0.003 {
		cl.Player.EntityPlayer.Motion.Z = 0
	}

	cl.Player.travel(maths.Vec3d{X: cl.Player.Controller.LeftImpulse, Z: cl.Player.Controller.ForwardImpulse})

	dX := cl.Player.EntityPlayer.Position.X - cl.Player.EntityPlayer.LastPosition.X
	dY := cl.Player.EntityPlayer.Position.Y - cl.Player.EntityPlayer.LastPosition.Y
	dZ := cl.Player.EntityPlayer.Position.Z - cl.Player.EntityPlayer.LastPosition.Z
	dPitch := cl.Player.EntityPlayer.Rotation.Y - cl.Player.EntityPlayer.LastRotation.Y
	dYaw := cl.Player.EntityPlayer.Rotation.X - cl.Player.EntityPlayer.LastRotation.X
	cl.Player.PositionReminder++
	flag = dX*dX+dY*dY+dZ*dZ > 4e-8 || cl.Player.PositionReminder >= 20
	flag2 := dPitch != 0 || dYaw != 0
	var err error
	if cl.Player.EntityPlayer.IsPassenger() {
		err = cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerPositionRotation, pk.Double(cl.Player.EntityPlayer.Position.X), pk.Double(-999), pk.Double(cl.Player.EntityPlayer.Position.Z), pk.Float(dYaw), pk.Float(dPitch), pk.Boolean(cl.Player.EntityPlayer.OnGround)))
		flag = false
	}

	if flag && flag2 {
		err = cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerPositionRotation, pk.Double(cl.Player.EntityPlayer.Position.X), pk.Double(cl.Player.EntityPlayer.Position.Y), pk.Double(cl.Player.EntityPlayer.Position.Z), pk.Float(dYaw), pk.Float(dPitch), pk.Boolean(cl.Player.EntityPlayer.OnGround)))
	}

	if flag {
		err = cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerPosition, pk.Double(cl.Player.EntityPlayer.Position.X), pk.Double(cl.Player.EntityPlayer.Position.Y), pk.Double(cl.Player.EntityPlayer.Position.Z), pk.Boolean(cl.Player.EntityPlayer.OnGround)))
		cl.Player.EntityPlayer.LastPosition = cl.Player.EntityPlayer.Position
		cl.Player.PositionReminder = 0
	}

	if flag2 {
		err = cl.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerRotation, pk.Float(dYaw), pk.Float(dPitch), pk.Boolean(cl.Player.EntityPlayer.OnGround)))
		cl.Player.EntityPlayer.LastRotation = cl.Player.EntityPlayer.Rotation
	}

	cl.Player.EntityPlayer.LastOnGround = cl.Player.EntityPlayer.OnGround
	return err
}

func (pl *Player) travel(position maths.Vec3d) {
	isSprinting := pl.Controller.Sprint
	feetBlock := pl.World.MustGetBlock(pl.EntityPlayer.Position)

	var gravity = enums.Gravity

	if pl.EntityPlayer.IsPotionActive(effects.SlowFalling) {
		gravity = enums.SlowFallGravity
	}

	switch feetBlock {
	case block.Water:
		if isSprinting {
			gravity = enums.WaterSprintDrag
		} else {
			gravity = enums.WaterDrag
		}

		// TODO: Depth strider
		speed := 0.02

		if pl.EntityPlayer.IsPotionActive(effects.DolphinsGrace) {
			gravity = enums.WaterDolphinDrag
		}

		pl.EntityPlayer.Motion.Add(inputVector(speed, pl.EntityPlayer.Rotation.Y, position))
		pl.move(enums.MoverTypeSelf, pl.EntityPlayer.Motion)
		pl.EntityPlayer.Motion.MulScalar(1, 0.3, 1)
	case block.Lava:
		pl.EntityPlayer.Motion.Add(inputVector(0.02, pl.EntityPlayer.Rotation.Y, position))
		pl.move(enums.MoverTypeSelf, pl.EntityPlayer.Motion)
		pl.EntityPlayer.Motion.MulScalar(0.5, 0.5, 0.5)
	default: // When air
		blockFriction := pl.World.MustGetBlockOnSide(pl.EntityPlayer.Position, maths.Down).Friction
		inertia := min(0.91, blockFriction*0.91)
		pl.EntityPlayer.Motion.Add(inputVector(pl.frictionInfluencedSpeed(inertia), pl.EntityPlayer.Rotation.Y, position))
		pl.move(enums.MoverTypeSelf, pl.EntityPlayer.Motion)
		motion := pl.EntityPlayer.Motion
		if pl.CollidedHorizontally || pl.Controller.Jump /*|| pl.EntityPlayer.IsOnLadder()*/ {
			motion.Y = 0.2
		}

		if pl.EntityPlayer.IsPotionActive(effects.Levitation) {
			gravity = (0.05*(float64(pl.EntityPlayer.GetPotionEffect(effects.Levitation).Amplifier)+1) - motion.Y) * 0.2
		}

		pl.EntityPlayer.Motion.MulScalar(inertia, 1, inertia)

		fmt.Println(gravity*0.98, pl.EntityPlayer.Motion.Y)
		if !pl.EntityPlayer.OnGround {
			pl.EntityPlayer.Motion.Y -= gravity * 0.98
		}
		//fmt.Println(pl.EntityPlayer.Motion.Y)
	}
}

func (pl *Player) move(mover enums.MoverType, motion maths.Vec3d) {
	/*feet, err := pl.World.GetBlock(pl.EntityPlayer.Position)
	if err != nil {
		return
	}*/

	if mover == enums.MoverTypePiston {
		motion = limitPistonMovement(motion)
	}

	//pl.backOffFromEdge(motion, mover)

	var xCollision, zCollision bool
	pl.CollidedHorizontally, xCollision, zCollision = pl.World.CollidesHorizontally(pl.EntityPlayer.BoundingBox)
	pl.CollidedVertically = pl.World.CollidesVertically(pl.EntityPlayer.BoundingBox)
	pl.CollidedVerticallyBelow = pl.CollidedVertically && pl.EntityPlayer.Motion.Y <= 0
	pl.EntityPlayer.OnGround = pl.CollidedVerticallyBelow

	if pl.CollidedHorizontally {
		pl.CollidedHorizontallyMinor = pl.isCollidedHorizontallyMinor(motion)
		var x, z float64
		if !xCollision {
			x = motion.X
		}

		if !zCollision {
			z = motion.Z
		}
		pl.EntityPlayer.Motion.SetScalar(x, motion.Y, z)
	}

	pl.EntityPlayer.Position.Add(motion)
}

func (pl *Player) collide(motion maths.Vec3d) {
	/*boundingBox := pl.EntityPlayer.BoundingBox
	boundingBox.Expand(motion.Spread())
	entities := pl.World.GetEntitiesInAABB(boundingBox)
	boundingBox.Unexpand(motion.Spread())*/
}

func inputVector(speed, yRot float64, pos maths.Vec3d) maths.Vec3d {
	length := pos.Length()
	if length < 1e-7 {
		return maths.Vec3d{}
	}

	vec := pos
	if length > 1 {
		vec.Normalize()
	}
	vec.Scale(speed)
	f := math.Sin(yRot * (math.Pi / 180))
	h := math.Cos(yRot * (math.Pi / 180))
	return maths.Vec3d{X: vec.X*h - vec.Z*f, Y: vec.Y, Z: vec.Z*h + vec.X*f}
}

func (pl *Player) frictionInfluencedSpeed(inertia float64) float64 {
	if !pl.EntityPlayer.OnGround {
		return pl.Abilities.FlyingSpeed
	}

	var speed = enums.PlayerSpeed
	if pl.Controller.Sprint {
		speed = enums.SprintSpeed
	}
	if pl.Controller.Sneak {
		speed = enums.SneakSpeed
	}
	return speed * (0.216 / (speed * speed * speed))
}

func (pl *Player) isCollidedHorizontallyMinor(motion maths.Vec3d) bool {
	return false
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

	box := pl.EntityPlayer.BoundingBox

	for d0 != 0 {
		box.Move(d0, -1, 0)
		if d0 < 0.05 && d0 >= -0.05 {
			d0 = 0
		} else if d0 > 0 {
			d0 -= 0.05
		} else {
			d0 += 0.05
		}
	}

	box = pl.EntityPlayer.BoundingBox

	for d1 != 0 {
		box.Move(0, -1, d1)
		if d1 < 0.05 && d1 >= -0.05 {
			d1 = 0
		} else if d1 > 0 {
			d1 -= 0.05
		} else {
			d1 += 0.05
		}
	}

	box = pl.EntityPlayer.BoundingBox

	for d0 != 0 && d1 != 0 {
		box.Move(d0, -1, d1)
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
