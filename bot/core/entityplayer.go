package core

import (
	"github.com/google/uuid"
)

type EntityPlayer struct {
	*EntityLiving
	Username string
}

func (e *EntityPlayer) GetName() string {
	return e.Type.Name
}

func (e *EntityPlayer) GetDisplayName() string {
	return e.Username // e.Username replaces the generic e.Type.DisplayName (which is "Player")
}

func NewEntityPlayer(displayName string, id int32, uuid uuid.UUID, t int32, x, y, z float64, yaw, pitch float64) *EntityPlayer {
	return &EntityPlayer{
		EntityLiving: NewEntityLiving(id, uuid, t, x, y, z, yaw, pitch),
		Username:     displayName,
	}
}
