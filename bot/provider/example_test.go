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
		return auth.Profile.Name == "aluwakbar"
	})

	//Login
	if err := c.JoinServer("ToxicNRV.aternos.me:33835"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler[Client]{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler[Client]{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler[Client]{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler[Client]{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler[Client]{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler[Client]{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler[Client]{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler[Client]{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler[Client]{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler[Client]{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler[Client]{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler[Client]{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler[Client]{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler[Client]{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler[Client]{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler[Client]{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler[Client]{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
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
		PacketHandler[Client]{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler[Client]{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler[Client]{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler[Client]{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler[Client]{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler[Client]{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler[Client]{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler[Client]{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler[Client]{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler[Client]{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler[Client]{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler[Client]{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler[Client]{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler[Client]{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler[Client]{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler[Client]{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler[Client]{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler[Client]{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler[Client]{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); err != nil {
		log.Fatal(err)
	}
}
