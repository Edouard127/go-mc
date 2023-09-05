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
	c := NewClient(microsoft.LoginFromCache(nil))

	//Login
	if err := c.JoinServer("2b2t.org"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler[Client]{ID: packetid.CPacketSetContainerContent, Priority: 50, F: SetContainerContent},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerSlot, Priority: 50, F: SetContainerSlot},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerProperty, Priority: 50, F: SetContainerProperty},

		/* Physic */
		PacketHandler[Client]{ID: packetid.CPacketChunkData, Priority: 50, F: ChunkData},
		PacketHandler[Client]{ID: packetid.CPacketUnloadChunk, Priority: 50, F: UnloadChunk},
		PacketHandler[Client]{ID: packetid.CPacketExplosion, Priority: 50, F: Explosion},

		/* Entities */
		PacketHandler[Client]{ID: packetid.CPacketSpawnEntity, Priority: 50, F: SpawnEntity},
		PacketHandler[Client]{ID: packetid.CPacketSpawnExperienceOrb, Priority: 50, F: SpawnExperienceOrb},
		PacketHandler[Client]{ID: packetid.CPacketSpawnPlayer, Priority: 50, F: SpawnPlayer},
		PacketHandler[Client]{ID: packetid.CPacketEntityAnimation, Priority: 50, F: EntityAnimation},
		PacketHandler[Client]{ID: packetid.CPacketBlockEntityData, Priority: 50, F: BlockEntityData},
		PacketHandler[Client]{ID: packetid.CPacketBlockAction, Priority: 50, F: BlockAction},
		PacketHandler[Client]{ID: packetid.CPacketBlockUpdate, Priority: 50, F: BlockChange},
		PacketHandler[Client]{ID: packetid.CPacketEntityPosition, Priority: 50, F: EntityPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityPositionRotation, Priority: 50, F: EntityPositionRotation},
		PacketHandler[Client]{ID: packetid.CPacketEntityRotation, Priority: 50, F: EntityRotation},
		PacketHandler[Client]{ID: packetid.CPacketVehicleMove, Priority: 50, F: VehicleMove},
		PacketHandler[Client]{ID: packetid.CPacketLookAt, Priority: 50, F: LookAt},
		PacketHandler[Client]{ID: packetid.CPacketPlayerPosition, Priority: 50, F: PlayerPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityEffect, Priority: 50, F: EntityEffect},
		PacketHandler[Client]{ID: packetid.CPacketSetEntityVelocity, Priority: 50, F: EntityVelocity},

		PacketHandler[Client]{ID: packetid.CPacketPlayerAbilities, Priority: 50, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); err != nil {
		log.Fatal(err)
	}
}

func TestExampleClient_JoinServer_offline(t *testing.T) {
	c := NewClient(data.DefaultAuth)

	//Login
	if err := c.JoinServer("127.50.50.1"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler[Client]{ID: packetid.CPacketSetContainerContent, Priority: 50, F: SetContainerContent},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerSlot, Priority: 50, F: SetContainerSlot},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerProperty, Priority: 50, F: SetContainerProperty},

		/* Physic */
		PacketHandler[Client]{ID: packetid.CPacketChunkData, Priority: 50, F: ChunkData},
		PacketHandler[Client]{ID: packetid.CPacketExplosion, Priority: 50, F: Explosion},

		/* Entities */
		PacketHandler[Client]{ID: packetid.CPacketSpawnEntity, Priority: 50, F: SpawnEntity},
		PacketHandler[Client]{ID: packetid.CPacketSpawnExperienceOrb, Priority: 50, F: SpawnExperienceOrb},
		PacketHandler[Client]{ID: packetid.CPacketSpawnPlayer, Priority: 50, F: SpawnPlayer},
		PacketHandler[Client]{ID: packetid.CPacketEntityAnimation, Priority: 50, F: EntityAnimation},
		PacketHandler[Client]{ID: packetid.CPacketBlockEntityData, Priority: 50, F: BlockEntityData},
		PacketHandler[Client]{ID: packetid.CPacketBlockAction, Priority: 50, F: BlockAction},
		PacketHandler[Client]{ID: packetid.CPacketBlockUpdate, Priority: 50, F: BlockChange},
		PacketHandler[Client]{ID: packetid.CPacketEntityPosition, Priority: 50, F: EntityPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityPositionRotation, Priority: 50, F: EntityPositionRotation},
		PacketHandler[Client]{ID: packetid.CPacketEntityRotation, Priority: 50, F: EntityRotation},
		PacketHandler[Client]{ID: packetid.CPacketVehicleMove, Priority: 50, F: VehicleMove},
		PacketHandler[Client]{ID: packetid.CPacketLookAt, Priority: 50, F: LookAt},
		PacketHandler[Client]{ID: packetid.CPacketPlayerPosition, Priority: 50, F: PlayerPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityEffect, Priority: 50, F: EntityEffect},
		PacketHandler[Client]{ID: packetid.CPacketSetEntityVelocity, Priority: 50, F: EntityVelocity},

		PacketHandler[Client]{ID: packetid.CPacketPlayerAbilities, Priority: 50, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); err != nil {
		log.Fatal(err)
	}
}
