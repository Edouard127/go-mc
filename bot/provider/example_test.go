package provider

import (
	"github.com/Edouard127/go-mc/auth/data"
	"github.com/Edouard127/go-mc/auth/microsoft"
	"github.com/Edouard127/go-mc/data/packetid"
	"log"
	"testing"
)

func TestExamplePingAndList(t *testing.T) {
	resp, delay, err := PingAndList("")
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}

	log.Println("Status:", resp)
	log.Println("Delay:", delay)
}

func TestExampleClient_JoinServer_online(t *testing.T) {
	c := NewClient()

	c.Auth = microsoft.LoginFromCache(func(auth data.Auth) bool {
		return auth.Name == "Kamigen"
	})

	//Login
	if err := c.JoinServer("localhost:25565"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); err != nil {
		log.Fatal(err)
	}
}

func TestExampleClient_JoinServer_offline(t *testing.T) {
	c := NewClient()

	//Login
	if err := c.JoinServer("127.0.0.1"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); err != nil {
		log.Fatal(err)
	}
}
