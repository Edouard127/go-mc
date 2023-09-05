package core

import (
	"github.com/Edouard127/go-mc/data/entity"
	"github.com/Edouard127/go-mc/maths"
	"github.com/google/uuid"
)

type UnaliveEntity struct {
	Type          entity.TypeEntity
	ID            int32
	UUID          uuid.UUID
	lastPosition  maths.Vec3d
	Position      maths.Vec3d
	EyePosition   maths.Vec3d
	Rotation      maths.Vec2d
	Motion        maths.Vec3d
	BoundingBox   maths.AxisAlignedBB[float64]
	Width, Height float64
	Vehicle       Entity
	Passengers    []Entity
	dataManager   map[int32]interface{}
}

type Entity interface {
	GetName() string
	GetDisplayName() string
	GetType() entity.TypeEntity
	GetID() int32
	GetUUID() uuid.UUID
	GetPosition() maths.Vec3d
	GetRotation() maths.Vec2d
	GetMotion() maths.Vec3d
	GetBoundingBox() maths.AxisAlignedBB[float64]
	GetWidth() float64
	GetHeight() float64
	GetDataManager() map[int32]interface{}
	IsLivingEntity() bool
	IsPlayer() bool
	SetPosition(x, y, z float64)
	SetRotation(yaw, pitch float64)
	SetMotion(x, y, z float64)
	SetSize(width, height float64)
	Equals(other entity.TypeEntity) bool
	IsPassenger() bool
	IsVehicle() bool
	GetPassengers() []Entity
}

func (e *UnaliveEntity) GetName() string {
	return e.Type.Name
}

func (e *UnaliveEntity) GetDisplayName() string {
	return e.Type.DisplayName
}

func (e *UnaliveEntity) GetType() entity.TypeEntity {
	return e.Type
}

func (e *UnaliveEntity) GetID() int32 {
	return e.ID
}

func (e *UnaliveEntity) GetUUID() uuid.UUID {
	return e.UUID
}

func (e *UnaliveEntity) GetPosition() maths.Vec3d {
	return e.Position
}

func (e *UnaliveEntity) GetRotation() maths.Vec2d {
	return e.Rotation
}

func (e *UnaliveEntity) GetMotion() maths.Vec3d {
	return e.Motion
}

func (e *UnaliveEntity) GetBoundingBox() maths.AxisAlignedBB[float64] {
	return e.BoundingBox

}

func (e *UnaliveEntity) GetWidth() float64 {
	return e.Width
}

func (e *UnaliveEntity) GetHeight() float64 {
	return e.Height
}

func (e *UnaliveEntity) GetDataManager() map[int32]interface{} {
	return e.dataManager
}

func (e *UnaliveEntity) IsLivingEntity() bool {
	return false
}

func (e *UnaliveEntity) IsPlayer() bool {
	return false
}

func (e *UnaliveEntity) SetPosition(x, y, z float64) {
	e.lastPosition = e.Position
	e.Position = maths.Vec3d{X: x, Y: y, Z: z}
	e.EyePosition = maths.Vec3d{X: x, Y: y + e.Height*0.85, Z: z}
}

func (e *UnaliveEntity) SetRotation(yaw, pitch float64) {
	e.Rotation = maths.Vec2d{X: yaw, Y: pitch}
}

func (e *UnaliveEntity) SetMotion(x, y, z float64) {
	e.Motion = maths.Vec3d{X: x, Y: y, Z: z}
}

func (e *UnaliveEntity) SetSize(width, height float64) {
	if width != e.Width || height != e.Height {
		f := e.Width
		e.Width = width
		e.Height = height

		if e.Width < f {
			d0 := width / 2.0
			e.BoundingBox = maths.AxisAlignedBB[float64]{
				MinX: e.Position.X - d0,
				MinY: e.Position.Y,
				MinZ: e.Position.Z - d0,
				MaxX: e.Position.X + d0,
				MaxY: e.Position.Y + height,
				MaxZ: e.Position.Z + d0,
			}
		}

		aabb := e.BoundingBox
		e.BoundingBox = maths.AxisAlignedBB[float64]{
			MinX: aabb.MinX,
			MinY: aabb.MinY,
			MinZ: aabb.MinZ,
			MaxX: aabb.MaxX + e.Width,
			MaxY: aabb.MaxY + e.Height,
			MaxZ: aabb.MaxZ + e.Width,
		}
	}
}

func (e *UnaliveEntity) Equals(other entity.TypeEntity) bool {
	return e.Type == other
}

func (e *UnaliveEntity) IsPassenger() bool {
	return e.Vehicle != nil
}

func (e *UnaliveEntity) IsVehicle() bool {
	return len(e.Passengers) > 0
}

func (e *UnaliveEntity) GetPassengers() []Entity {
	return e.Passengers
}

func NewEntity(id int32, uuid uuid.UUID, t int32, x, y, z float64, yaw, pitch float64) *UnaliveEntity {
	entityType := entity.TypeEntityByID[t]
	e := &UnaliveEntity{
		Type:     entityType,
		ID:       id,
		UUID:     uuid,
		Position: maths.Vec3d{X: x, Y: y, Z: z},
		Rotation: maths.Vec2d{X: yaw, Y: pitch},
	}
	e.SetSize(entityType.Width, entityType.Height)
	return e
}
