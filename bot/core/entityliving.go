package core

import (
	"github.com/Edouard127/go-mc/data/effects"
	"github.com/Edouard127/go-mc/data/enums"
	"github.com/Edouard127/go-mc/data/item"
	"github.com/Edouard127/go-mc/maths"
	"github.com/google/uuid"
)

var EyePosVec = maths.Vec3d[float64]{Y: 1.62}
var EyePos = 1.62

type EntityLiving struct {
	*Entity
	health                  float32
	minHealth               float32
	maxHealth               float32
	Food                    int32
	MaxFood                 int32
	Saturation              float32
	Absorption              float32
	ActiveItem              *item.Item
	ActiveItemStackUseCount int32
	ActivePotionEffects     map[int32]*effects.EffectStatus
	dead                    bool
	OnGround                bool
	MoveStrafing            float32
	MoveForward             float32
	MoveVertical            float32
}

type EntityLivingInterface interface {
	EntityInterface
	GetHealth(absorption bool) float32
	SetHealth(health float32) bool
	GetEyePos() maths.Vec3d[float64]
	IsDead() bool
	IsPotionActive(effect effects.Effect) bool
	GetPotionEffect(effect effects.Effect) *effects.EffectStatus
}

func (e *EntityLiving) GetName() string {
	return e.Type.Name
}

func (e *EntityLiving) GetDisplayName() string {
	return e.Type.DisplayName
}

func (e *EntityLiving) GetHealth(absorption bool) float32 {
	if absorption {
		return e.health + e.Absorption
	}
	return e.health
}

func (e *EntityLiving) SetHealth(health float32) bool {
	e.health = health
	if e.IsDead() {
		return true
	}
	return false
}

func (e *EntityLiving) GetEyePos() maths.Vec3d[float64] {
	return e.Position.Add(EyePosVec)
}

func (e *EntityLiving) IsDead() bool {
	return e.health <= e.minHealth
}

func (e *EntityLiving) IsPotionActive(effect effects.Effect) bool {
	_, ok := e.ActivePotionEffects[effect.ID]
	return ok
}

func (e *EntityLiving) GetPotionEffect(effect effects.Effect) *effects.EffectStatus {
	return e.ActivePotionEffects[effect.ID]
}

func (e *EntityLiving) IsInvulnerableTo(source enums.DamageSource) bool {
	return e.Entity.IsInvulnerableTo(source)
}

func (e *EntityLiving) IsLivingEntity() bool {
	return true
}

func NewEntityLiving(id int32, uuid uuid.UUID, t int32, x, y, z float64, yaw, pitch float64) *EntityLiving {
	return &EntityLiving{
		Entity:              NewEntity(id, uuid, t, x, y, z, yaw, pitch),
		ActivePotionEffects: make(map[int32]*effects.EffectStatus),
	}
}
