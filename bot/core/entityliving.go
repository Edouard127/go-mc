package core

import (
	"github.com/Edouard127/go-mc/data/effects"
	"github.com/Edouard127/go-mc/data/item"
	"github.com/Edouard127/go-mc/maths"
	"github.com/google/uuid"
)

var EyePosVec = maths.Vec3d{Y: 1.62}
var EyePos = 1.62

type EntityLiving struct {
	*UnaliveEntity
	Health                  float32
	MinHealth               float32
	MaxHealth               float32
	Food                    int32
	MaxFood                 int32
	Saturation              float32
	Absorption              float32
	ActiveItem              *item.Item
	ActiveItemStackUseCount int32
	ActivePotionEffects     map[int32]effects.EffectStatus
	OnGround                bool
	LastOnGround            bool
	MoveStrafing            float32
	MoveForward             float32
	MoveVertical            float32
}

func (e *EntityLiving) GetName() string {
	return e.Type.Name
}

func (e *EntityLiving) GetDisplayName() string {
	return e.Type.DisplayName
}

func (e *EntityLiving) GetHealth(absorption bool) float32 {
	if absorption {
		return e.Health + e.Absorption
	}
	return e.Health
}

func (e *EntityLiving) SetHealth(health float32) bool {
	e.Health = health
	return e.IsDead()
}

func (e *EntityLiving) GetEyePos() maths.Vec3d {
	c := e.Position
	c.Y += EyePos
	return c
}

func (e *EntityLiving) IsDead() bool {
	return e.Health <= e.MinHealth
}

func (e *EntityLiving) IsPotionActive(effect effects.Effect) bool {
	_, ok := e.ActivePotionEffects[effect.ID]
	return ok
}

func (e *EntityLiving) GetPotionEffect(effect effects.Effect) effects.EffectStatus {
	return e.ActivePotionEffects[effect.ID]
}

func (e *EntityLiving) IsLivingEntity() bool {
	return true
}

func NewEntityLiving(id int32, uuid uuid.UUID, t int32, x, y, z float64, yaw, pitch float64) *EntityLiving {
	return &EntityLiving{
		UnaliveEntity:       NewEntity(id, uuid, t, x, y, z, yaw, pitch),
		ActivePotionEffects: make(map[int32]effects.EffectStatus),
	}
}
