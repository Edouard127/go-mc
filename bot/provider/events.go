package provider

import (
	"context"
	"fmt"
	"github.com/Edouard127/go-mc/bot/core"
	"github.com/Edouard127/go-mc/bot/world"
	"github.com/Edouard127/go-mc/data/effects"
	"github.com/Edouard127/go-mc/data/entity"
	"github.com/Edouard127/go-mc/data/grids"
	"github.com/Edouard127/go-mc/level"
	"unsafe"

	"github.com/google/uuid"

	"github.com/Edouard127/go-mc/chat"
	"github.com/Edouard127/go-mc/data/packetid"
	. "github.com/Edouard127/go-mc/data/slots"
	pk "github.com/Edouard127/go-mc/net/packet"
)

// Attach attaches the core handlers to the client
func Attach(c *Client) *Client {
	c.Events.AddListener(
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketLogin, F: JoinGame},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketKeepAlive, F: KeepAlive},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketChatMessage, F: ChatMessage},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketSystemMessage, F: ChatMessage},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketDisconnect, F: Disconnect},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketSetHealth, F: UpdateHealth},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketUpdateTime, F: TimeUpdate},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketPlayerInfoUpdate, F: PlayerInfoUpdate},
		PacketHandler[Client]{Priority: 50, ID: packetid.CPacketPlayerInfoRemove, F: PlayerInfoUpdate},
	)

	c.Events.AddTicker(
		TickHandler[Client]{Priority: 50, F: Step},
		TickHandler[Client]{Priority: 50, F: RunTransactions},
	)

	return c
}

func SpawnEntity(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id                  pk.VarInt
		euuid               pk.UUID
		t                   pk.Byte
		x, y, z             pk.Double
		pitch, yaw, headYaw pk.Angle
		data                pk.VarInt
		vX, vY, vZ          pk.Short
	)

	if err := p.Scan(&id, &euuid, &t, &x, &y, &z, &pitch, &yaw, &headYaw, &data, &vX, &vY, &vZ); err != nil {
		return fmt.Errorf("unable to read SpawnEntity packet: %w", err)
	}

	if entity.TypeEntityByID[int32(t)].IsLiving() {
		c.World.Add(core.NewEntityLiving(int32(id), uuid.UUID(euuid), int32(t), float64(x), float64(y), float64(z), float64(yaw), float64(pitch)))
	} else {
		c.World.Add(core.NewEntity(int32(id), uuid.UUID(euuid), int32(t), float64(x), float64(y), float64(z), float64(yaw), float64(pitch)))
	}

	return nil
}

func SpawnExperienceOrb(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID pk.VarInt
		x, y, z  pk.Double
		count    pk.Short
	)

	if err := p.Scan(&entityID, &x, &y, &z, &count); err != nil {
		return fmt.Errorf("unable to read SpawnExperienceOrb packet: %w", err)
	}

	return nil
}

func SpawnPlayer(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id         pk.VarInt
		euuid      pk.UUID
		x, y, z    pk.Double
		yaw, pitch pk.Angle
	)

	if err := p.Scan(&id, &euuid, &x, &y, &z, &yaw, &pitch); err != nil {
		return fmt.Errorf("unable to read SpawnPlayer packet: %w", err)
	}

	c.World.Add(core.NewEntityPlayer(c.PlayerList.GetPlayer(uuid.UUID(euuid)).Name, int32(id), uuid.UUID(euuid), 116, float64(x), float64(y), float64(z), float64(yaw), float64(pitch)))

	return nil
}

func EntityAnimation(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID  pk.VarInt
		animation pk.Byte
	)

	if err := p.Scan(&entityID, &animation); err != nil {
		return fmt.Errorf("unable to read Animation packet: %w", err)
	}

	return nil
}

func AwardStatistics(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var count pk.VarInt
	var statistics []struct {
		Name  pk.String
		Value pk.VarInt
	}*/

	/*if err := p.Scan(&count, &statistics); err != nil {
		return err
	}*/

	return nil
}

func SetBlockDestroyStage(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID pk.VarInt
		location pk.Position
		stage    pk.Byte
	)

	if err := p.Scan(&entityID, &location, &stage); err != nil {
		return fmt.Errorf("unable to read BlockBreakAnimation packet: %w", err)
	}

	return nil
}

func BlockEntityData(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var location pk.Position
	var action pk.Byte
	var nbtData pk

	if err := p.Scan(&location, &action, &nbtData); err != nil {
		return err
	}*/

	return nil
}

func BlockAction(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		location    pk.Position
		actionID    pk.Byte
		actionParam pk.Byte
		blockType   pk.VarInt
	)

	if err := p.Scan(&location, &actionID, &actionParam, &blockType); err != nil {
		return fmt.Errorf("unable to read BlockAction packet: %w", err)
	}

	return nil
}

func BlockChange(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		location  pk.Position
		blockType pk.VarInt
	)

	if err := p.Scan(&location, &blockType); err != nil {
		return fmt.Errorf("unable to read BlockChange packet: %w", err)
	}

	return nil
}

func BossBar(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var uuid pk.UUID
	var action pk.Byte

	var err error

	if err = p.Scan(&uuid, &action); err != nil {
		return fmt.Errorf("unable to read BossBar packet: %w", err)
	}

	switch action {
	case 0:
		var (
			title     pk.String
			health    pk.Float
			color     pk.Byte
			divisions pk.Byte
			flags     pk.Byte
		)

		if err = p.Scan(&title, &health, &color, &divisions, &flags); err != nil {
			return fmt.Errorf("unable to read BossBar packet: %w", err)
		}

	case 1:
		var health pk.Float

		err = p.Scan(&health)
	case 2:
		var title pk.String

		err = p.Scan(&title)
	case 3:
		var color pk.Byte

		err = p.Scan(&color)
	case 4:
		var division pk.Byte

		err = p.Scan(&division)
	case 5:
		var flags pk.Byte

		err = p.Scan(&flags)
	case 6:
		//
	}

	return nil
}

func ServerDifficulty(c *Client, p pk.Packet) error {
	var difficulty pk.Byte

	return p.Scan(&difficulty)
}

func TabComplete(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var count pk.VarInt
	var matches []pk.String

	if err := p.Scan(&count, &matches); err != nil {
		return err
	}*/

	return nil
}

func ChatMessage(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var message chat.Message

	if err := p.Scan(&message); err != nil {
		return fmt.Errorf("unable to read ChatMessage packet: %w", err)
	}

	fmt.Println(message)

	return nil
}

func MultiBlockChange(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var chunkX pk.Int
	var chunkZ pk.Int
	var recordCount pk.VarInt
	var records []struct {
		Position pk.Byte
		BlockID  pk.VarInt
	}

	if err := p.Scan(&chunkX, &chunkZ, &recordCount, &records); err != nil {
		return err
	}*/

	return nil
}

func SetContainerContent(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id    pk.UnsignedByte
		state pk.VarInt
		data  []Slot
		item  Slot
		err   error
	)

	if err := p.Scan(&id, &state, pk.Array(&data), &item); err != nil {
		return fmt.Errorf("failed to scan SetContainerContent: %w", err)
	}

	container := grids.Containers[int(id)]

	for i := range data {
		err = container.SetSlot(i, &data[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func CloseContainer(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return fmt.Errorf("unable to read CloseContainer packet: %w", err)
	}

	return nil
}

func CloseWindow(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return fmt.Errorf("unable to read CloseWindow packet: %w", err)
	}

	return nil
}

func SetContainerProperty(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		windowID pk.Byte
		property pk.Short
		value    pk.Short
	)

	if err := p.Scan(&windowID, &property, &value); err != nil {
		return fmt.Errorf("unable to read WindowProperty packet: %w", err)
	}

	return nil
}

func SetContainerSlot(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		containerId pk.Byte
		stateId     pk.VarInt
		slotId      pk.Short
		data        Slot
	)
	if err := p.Scan(&containerId, &stateId, &slotId, &data); err != nil {
		return fmt.Errorf("failed to scan SetSlot: %w", err)
	}

	c.Player.Manager.StateID = int32(stateId)

	if containerId == -1 && slotId == -1 {
		c.Player.Manager.HeldItem = data.Item()
		return nil
	}

	return grids.Containers[int(containerId)].SetSlot(int(slotId), &data)
}

func SetCooldown(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		itemID pk.VarInt
		ticks  pk.VarInt
	)

	if err := p.Scan(&itemID, &ticks); err != nil {
		return fmt.Errorf("unable to read SetCooldown packet: %w", err)
	}

	return nil
}

func PluginMessage(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		channel pk.String
		data    pk.ByteArray
	)

	if err := p.Scan(&channel, &data); err != nil {
		return fmt.Errorf("unable to read PluginMessage packet: %w", err)
	}

	return nil
}

func NamedSoundEffect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		soundName      pk.String
		soundCategory  pk.Byte
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundName, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return fmt.Errorf("unable to read NamedSoundEffect packet: %w", err)
	}

	return nil
}

func Disconnect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var reason chat.Message
	if err := p.Scan(&reason); err != nil {
		return fmt.Errorf("failed to scan Disconnect: %w", err)
	}

	return nil
}

func EntityStatus(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var entityID pk.Int
	var entityStatus pk.Byte

	if err := p.Scan(&entityID, &entityStatus); err != nil {
		return fmt.Errorf("unable to read EntityStatus packet: %w", err)
	}

	return nil
}

func Explosion(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		x, y, z    pk.Float
		radius     pk.Float
		records    pk.VarInt
		data       = make([][3]pk.VarInt, 0)
		mX, mY, mZ pk.Float
	)

	if err := p.Scan(&x, &y, &z, &radius, &records); err != nil {
		return fmt.Errorf("unable to read Explosion packet: %w", err)
	}

	data = make([][3]pk.VarInt, records)
	for i := pk.VarInt(0); i < records; i++ {
		if err := p.Scan(&data[i][0], &data[i][1], &data[i][2]); err != nil {
			return fmt.Errorf("unable to read Explosion packet: %w", err)
		}
	}

	if err := p.Scan(&mX, &mY, &mZ); err != nil {
		return fmt.Errorf("unable to read Explosion packet: %w", err)
	}

	return nil
}

func UnloadChunk(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var chunk level.ChunkPos

	if err := p.Scan(&chunk); err != nil {
		return fmt.Errorf("unable to read UnloadChunk packet: %w", err)
	}

	delete(c.World.Columns, chunk)
	return nil
}

func ChangeGameState(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var reason pk.UnsignedByte
	var value pk.Float

	if err := p.Scan(&reason, &value); err != nil {
		return fmt.Errorf("unable to read ChangeGameState packet: %w", err)
	}

	return nil
}

func KeepAlive(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	p.ID = packetid.SPacketKeepAlive
	err := c.Conn.WritePacket(p)
	if err != nil {
		return fmt.Errorf("unable to write KeepAlive packet: %w", err)
	}

	return nil
}

func ChunkData(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		ChunkPos level.ChunkPos
		Chunk    level.Chunk
	)

	if err := p.Scan(
		&ChunkPos, &Chunk,
	); err != nil {
		return fmt.Errorf("unable to read ChunkData packet: %w", err)
	}

	c.World.Columns[ChunkPos] = &Chunk

	return nil
}

func Effect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var effectID pk.Int
	var location pk.Position
	var data pk.Int
	var disableRelativeVolume pk.Boolean

	if err := p.Scan(&effectID, &location, &data, &disableRelativeVolume); err != nil {
		return fmt.Errorf("unable to read Effect packet: %w", err)
	}

	return nil
}

func Particle(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var particleID pk.String
	var longDistance pk.Boolean
	var x pk.Float
	var y pk.Float
	var z pk.Float
	var offsetX pk.Float
	var offsetY pk.Float
	var offsetZ pk.Float
	var particleData pk.Float
	var particleCount pk.Int
	var data []pk.Int

	if err := p.Scan(&particleID, &longDistance, &x, &y, &z, &offsetX, &offsetY, &offsetZ, &particleData, &particleCount, &data); err != nil {
		return err
	}*/

	return nil
}

func JoinGame(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	if err := p.Scan(
		(*pk.Int)(&c.Player.EntityPlayer.ID),
		(*pk.Boolean)(&c.Player.PlayerInfo.Hardcore),
		(*pk.UnsignedByte)(&c.Player.PlayerInfo.Gamemode),
		(*pk.Byte)(&c.Player.PlayerInfo.PrevGamemode),
		pk.Array((*[]pk.Identifier)(unsafe.Pointer(&c.Player.WorldInfo.DimensionNames))),
		pk.NBT(&c.Player.WorldInfo.DimensionCodec),
		(*pk.Identifier)(&c.Player.WorldInfo.DimensionType),
		(*pk.Identifier)(&c.Player.WorldInfo.DimensionName),
		(*pk.Long)(&c.Player.WorldInfo.HashedSeed),
		(*pk.VarInt)(&c.Player.WorldInfo.MaxPlayers),
		(*pk.VarInt)(&c.Player.WorldInfo.ViewDistance),
		(*pk.VarInt)(&c.Player.WorldInfo.SimulationDistance),
		(*pk.Boolean)(&c.Player.WorldInfo.ReducedDebugInfo),
		(*pk.Boolean)(&c.Player.WorldInfo.EnableRespawnScreen),
		(*pk.Boolean)(&c.Player.WorldInfo.IsDebug),
		(*pk.Boolean)(&c.Player.WorldInfo.IsFlat),
		pk.Optional[pk.Tuple, *pk.Tuple]{
			Has: &c.Player.WorldInfo.HasDeathLocation,
			Value: pk.Tuple{
				(*pk.Identifier)(&c.Player.WorldInfo.DimensionName),
				(*pk.Position)(&c.Player.WorldInfo.DeathPosition),
			},
		},
		(*pk.VarInt)(&c.Player.WorldInfo.PortalCooldown),
	); err != nil {
		return fmt.Errorf("unable to read JoinGame packet: %w", err)
	}

	if err := c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		packetid.SPacketPluginMessage,
		pk.Identifier("minecraft:brand"),
		pk.String(c.Player.Settings.Brand),
	)); err != nil {
		return fmt.Errorf("unable to write PluginMessage packet: %w", err)
	}

	if err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientSettings,
		pk.String(c.Player.Settings.Locale),
		pk.Byte(c.Player.Settings.ViewDistance),
		pk.VarInt(c.Player.Settings.ChatMode),
		pk.Boolean(c.Player.Settings.ChatColors),
		pk.UnsignedByte(c.Player.Settings.DisplayedSkinParts),
		pk.VarInt(c.Player.Settings.MainHand),
		pk.Boolean(c.Player.Settings.EnableTextFiltering),
		pk.Boolean(c.Player.Settings.AllowListing),
	)); err != nil {
		return fmt.Errorf("unable to write ClientSettings packet: %w", err)
	}

	c.World.Add(c.Player.EntityPlayer)

	return nil
}

func Map(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var Map world.Map

	if err := p.Scan(&Map); err != nil {
		return fmt.Errorf("unable to read Map packet: %w", err)
	}

	return nil
}

func Entity(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var entityID pk.Int

	if err := p.Scan(&entityID); err != nil {
		return fmt.Errorf("unable to read UnaliveEntity packet: %w", err)
	}

	return nil
}

func EntityPosition(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id         pk.VarInt
		dX, dY, dZ pk.Short
		onGround   pk.Boolean
	)

	if err := p.Scan(&id, &dX, &dY, &dZ, &onGround); err != nil {
		return fmt.Errorf("unable to read EntityPosition packet: %w", err)
	}

	if e := c.World.GetEntity(int32(id)); e != nil {
		nX := e.GetPosition().X + (float64(dX) / 4096)
		nY := e.GetPosition().Y + (float64(dY) / 4096)
		nZ := e.GetPosition().Z + (float64(dZ) / 4096)
		e.SetPosition(nX, nY, nZ)
	}

	return nil
}

func EntityPositionRotation(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id         pk.VarInt
		dX, dY, dZ pk.Short
		yaw, pitch pk.Angle
		onGround   pk.Boolean
	)

	if err := p.Scan(&id, &dX, &dY, &dZ, &yaw, &pitch, &onGround); err != nil {
		return fmt.Errorf("unable to read EntityPositionRotation packet: %w", err)
	}

	if e := c.World.GetEntity(int32(id)); e != nil {
		nX := e.GetPosition().X + (float64(dX) / 4096)
		nY := e.GetPosition().Y + (float64(dY) / 4096)
		nZ := e.GetPosition().Z + (float64(dZ) / 4096)
		e.SetPosition(nX, nY, nZ)
		e.SetRotation(float64(yaw), float64(pitch))
	}

	return nil
}

func EntityHeadRotation(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		EntityID pk.VarInt
		HeadYaw  pk.Angle
	)

	if err := p.Scan(&EntityID, &HeadYaw); err != nil {
		return fmt.Errorf("unable to read EntityHeadRotation packet: %w", err)
	}

	return nil
}

func EntityRotation(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		id         pk.VarInt
		yaw, pitch pk.Angle
		onGround   pk.Boolean
	)

	if err := p.Scan(&id, &yaw, &pitch, &onGround); err != nil {
		return fmt.Errorf("unable to read EntityRotation packet: %w", err)
	}

	if e := c.World.GetEntity(int32(id)); e != nil {
		e.SetRotation(float64(yaw), float64(pitch))
	}

	return nil
}

func VehicleMove(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return fmt.Errorf("unable to read VehicleMove packet: %w", err)
	}

	return nil
}

func OpenSignEditor(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return fmt.Errorf("unable to read OpenSignEditor packet: %w", err)
	}

	return nil
}

func CraftRecipeResponse(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var windowID pk.UnsignedByte
	var recipe pk.String

	if err := p.Scan(&windowID, &recipe); err != nil {
		return fmt.Errorf("unable to read CraftRecipeResponse packet: %w", err)
	}

	return nil
}

func PlayerAbilities(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var flags pk.UnsignedByte
	var flyingSpeed pk.Float
	var fov pk.Float

	if err := p.Scan(&flags, &flyingSpeed, &fov); err != nil {
		return fmt.Errorf("unable to read PlayerAbilities packet: %w", err)
	}

	return nil
}

func CombatEvent(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var event pk.Byte
	var duration pk.Int
	var entityID pk.Int
	var message pk.String

	if err := p.Scan(&event, &duration, &entityID, &message); err != nil {
		return fmt.Errorf("unable to read CombatEvent packet: %w", err)
	}

	return nil
}

func PlayerInfoRemove(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var players []pk.UUID

	if err := p.Scan(pk.Array(&players)); err != nil {
		return fmt.Errorf("unable to read PlayerInfoRemove packet: %w", err)
	}

	c.PlayerList.RemovePlayers(players)
	return nil
}

// Since Mojang has done an horrible job at designing this, I will simply ignore it
// until I need it, fuck off
func PlayerInfoUpdate(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*if err := p.Scan(c.PlayerList); err != nil {
		return fmt.Errorf("unable to read PlayerInfoUpdate packet: %w", err)
	}*/
	return nil
}

func LookAt(client *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		target       pk.VarInt
		x, y, z      pk.Double
		isEntity     pk.Boolean
		id           pk.VarInt
		entityTarget pk.VarInt
	)

	if err := p.Scan(&target, &x, &y, &z, &isEntity,
		pk.Optional[pk.Tuple, *pk.Tuple]{
			Has: &isEntity,
			Value: pk.Tuple{
				&id, &entityTarget,
			},
		}); err != nil {
		return fmt.Errorf("unable to read LookAt p: %w", err)
	}

	return nil
}

func SynchronizePlayerPosition(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		x, y, z    pk.Double
		yaw, pitch pk.Float
		flags      pk.Byte
		teleportID pk.VarInt
	)

	if err := p.Scan(&x, &y, &z, &yaw, &pitch, &flags, &teleportID); err != nil {
		return fmt.Errorf("unable to read SynchronizePlayerPosition packet: %w", err)
	}

	if flags&0x01 != 0 {
		c.Player.EntityPlayer.SetRelativePosition(float64(x), float64(y), float64(z))
		c.Player.EntityPlayer.SetRelativeRotation(float64(yaw), float64(pitch))
	} else {
		c.Player.EntityPlayer.SetPosition(float64(x), float64(y), float64(z))
		c.Player.EntityPlayer.SetRotation(float64(yaw), float64(pitch))
	}

	var err error
	if err = c.Conn.WritePacket(pk.Marshal(packetid.SPacketTeleportConfirm, teleportID)); err != nil {
		return fmt.Errorf("unable to write TeleportConfirm packet: %w", err)
	}

	if err = c.Conn.WritePacket(pk.Marshal(packetid.SPacketPlayerPositionRotation, x, y, z, yaw, pitch, pk.Boolean(false))); err != nil {
		return fmt.Errorf("unable to write PlayerPositionRotation packet: %w", err)
	}

	return nil
}

func UnlockRecipes(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var action pk.Byte
	var craftingBookOpen pk.Boolean
	var filter pk.Boolean
	var recipes []pk.String

	if err := p.Scan(&action, &craftingBookOpen, &filter, &recipes); err != nil {
		return err
	}*/

	return nil
}

func DestroyEntities(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var entityIDs []pk.Int

	if err := p.Scan(&entityIDs); err != nil {
		return err
	}*/

	return nil
}

func RemoveEntityEffect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var entityID pk.Int
	var effectID pk.Byte

	if err := p.Scan(&entityID, &effectID); err != nil {
		return fmt.Errorf("unable to read RemoveEntityEffect packet: %w", err)
	}

	return nil
}

func ResourcePackSend(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		url  pk.String
		hash pk.String
	)

	if err := p.Scan(&url, &hash); err != nil {
		return fmt.Errorf("unable to read ResourcePackSend packet: %w", err)
	}

	return nil
}

func Respawn(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var copyMeta bool
	if err := p.Scan(
		(*pk.String)(&c.Player.WorldInfo.DimensionType),
		(*pk.Identifier)(&c.Player.WorldInfo.DimensionName),
		(*pk.Long)(&c.Player.WorldInfo.HashedSeed),
		(*pk.UnsignedByte)(&c.Player.PlayerInfo.Gamemode),
		(*pk.Byte)(&c.Player.PlayerInfo.PrevGamemode),
		(*pk.Boolean)(&c.Player.WorldInfo.IsDebug),
		(*pk.Boolean)(&c.Player.WorldInfo.IsFlat),
		(*pk.Boolean)(&copyMeta),
	); err != nil {
		return fmt.Errorf("unable to read Respawn packet: %w", err)
	}

	return nil
}

func SelectAdvancementTab(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var hasID pk.Boolean
	var identifier pk.String

	if err := p.Scan(&hasID, &identifier); err != nil {
		return fmt.Errorf("unable to read SelectAdvancementTab packet: %w", err)
	}

	return nil
}

func InitializeBorder(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		action         pk.Byte
		radius         pk.Double
		oldRadius      pk.Double
		speed          pk.VarInt
		x              pk.Int
		z              pk.Int
		portalBoundary pk.Int
		warningTime    pk.VarInt
		warningBlocks  pk.VarInt
	)

	if err := p.Scan(&action, &radius, &oldRadius, &speed, &x, &z, &portalBoundary, &warningTime, &warningBlocks); err != nil {
		return fmt.Errorf("unable to read WorldBorder packet: %w", err)
	}

	return nil
}

func Camera(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var cameraID pk.Int

	if err := p.Scan(&cameraID); err != nil {
		return fmt.Errorf("unable to read Camera packet: %w", err)
	}

	return nil
}

func SetHeldItem(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var slot pk.Short

	if err := p.Scan(&slot); err != nil {
		return fmt.Errorf("unable to read SetHeldItem packet: %w", err)
	}

	c.Player.Manager.HeldItem = c.Player.Manager.Inventory.GetItem(int(slot))

	return nil
}

func DisplayScoreboard(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var position pk.Byte
	var name pk.String

	if err := p.Scan(&position, &name); err != nil {
		return fmt.Errorf("unable to read DisplayScoreboard packet: %w", err)
	}

	return nil
}

func AttachEntity(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID  pk.Int
		vehicleID pk.Int
		leash     pk.Boolean
	)

	if err := p.Scan(&entityID, &vehicleID, &leash); err != nil {
		return fmt.Errorf("unable to read AttachEntity packet: %w", err)
	}

	return nil
}

func EntityVelocity(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID                        pk.VarInt
		velocityX, velocityY, velocityZ pk.Short
	)

	if err := p.Scan(&entityID, &velocityX, &velocityY, &velocityZ); err != nil {
		return fmt.Errorf("unable to read EntityVelocity packet: %w", err)
	}

	if e := c.World.GetEntity(int32(entityID)); e != nil {
		e.SetMotion(float64(velocityX)/8000, float64(velocityY)/8000, float64(velocityZ)/8000)
		/*if c.Player.EntityPlayer.ID == int32(entityID) {
			c.Player.EntityPlayer.Motion.Y = 0.42
		}*/
	}

	return nil
}

func EntityEquipment(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var entityID pk.Int
	var slot pk.Short
	var item pk.Slot

	if err := p.Scan(&entityID, &slot, &item); err != nil {
		return err
	}*/

	return nil
}

func SetExperience(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		experienceBar   pk.Float
		levelInt        pk.VarInt
		totalExperience pk.VarInt
	)

	if err := p.Scan(&experienceBar, &levelInt, &totalExperience); err != nil {
		return fmt.Errorf("unable to read SetExperience packet: %w", err)
	}

	c.Player.ExpBar = float64(experienceBar)
	c.Player.Level = int32(levelInt)
	c.Player.TotalExp = int32(totalExperience)

	return nil
}

func UpdateHealth(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		health         pk.Float
		food           pk.VarInt
		foodSaturation pk.Float
	)

	if err := p.Scan(&health, &food, &foodSaturation); err != nil {
		return fmt.Errorf("unable to read UpdateHealth packet: %w", err)
	}

	if respawn := c.Player.EntityPlayer.SetHealth(float32(health)); respawn {
		if err := c.Player.Respawn(c); err != nil {
			return nil
		}
	}
	return nil
}

func ScoreboardObjective(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		name           pk.String
		mode           pk.Byte
		objectiveName  pk.String
		objectiveValue pk.String
		type_          pk.Byte
	)

	if err := p.Scan(&name, &mode, &objectiveName, &objectiveValue, &type_); err != nil {
		return fmt.Errorf("unable to read ScoreboardObjective packet: %w", err)
	}

	return nil
}

func SetPassengers(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var entityID pk.Int
	var passengerCount pk.VarInt
	var passengers pk.Ary[]pk.VarInt

	if err := p.Scan(&entityID, &passengers); err != nil {
		return err
	}*/

	return nil
}

func Teams(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var name pk.String
	var mode pk.Byte
	var teamName pk.String
	var displayName pk.String
	var prefix pk.String
	var suffix pk.String
	var friendlyFire pk.Byte
	var nameTagVisibility pk.String
	var collisionRule pk.String
	var color pk.Byte
	var playerCount pk.VarInt
	var players pk.Ary[]pk.String

	if err := p.Scan(&name, &mode, &teamName, &displayName, &prefix, &suffix, &friendlyFire, &nameTagVisibility, &collisionRule, &color, &players); err != nil {
		return err
	}*/

	return nil
}

func UpdateScore(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var name pk.String
	var action pk.Byte
	var objectiveName pk.String
	var value pk.VarInt

	if err := p.Scan(&name, &action, &objectiveName, &value); err != nil {
		return fmt.Errorf("unable to read UpdateScore packet: %w", err)
	}

	return nil
}

func SpawnPosition(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return fmt.Errorf("unable to read SpawnPosition packet: %w", err)
	}

	return nil
}

func TimeUpdate(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		WorldAge  pk.Long
		TimeOfDay pk.Long
	)

	if err := p.Scan(&WorldAge, &TimeOfDay); err != nil {
		return fmt.Errorf("unable to read TimeUpdate packet: %w", err)
	}
	return nil
}

func Title(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		action    pk.Byte
		fadeIn    pk.Int
		stay      pk.Int
		fadeOut   pk.Int
		title     pk.String
		subtitle  pk.String
		actionBar pk.String
	)

	if err := p.Scan(&action, &fadeIn, &stay, &fadeOut, &title, &subtitle, &actionBar); err != nil {
		return fmt.Errorf("unable to read Title packet: %w", err)
	}

	return nil
}

func SoundEffect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		soundID        pk.VarInt
		soundCategory  pk.VarInt
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundID, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return fmt.Errorf("unable to read SoundEffect packet: %w", err)
	}

	return nil
}

func PlayerListHeaderAndFooter(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		header pk.String
		footer pk.String
	)

	if err := p.Scan(&header, &footer); err != nil {
		return fmt.Errorf("unable to read PlayerListHeaderAndFooter packet: %w", err)
	}

	return nil
}

func CollectItem(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		collectedEntityID pk.Int
		collectorEntityID pk.Int
		pickupCount       pk.Int
	)

	if err := p.Scan(&collectedEntityID, &collectorEntityID, &pickupCount); err != nil {
		return fmt.Errorf("unable to read CollectItem packet: %w", err)
	}

	return nil
}

func EntityTeleport(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var entityID pk.VarInt
	var x pk.Double
	var y pk.Double
	var z pk.Double
	var yaw pk.Byte
	var pitch pk.Byte
	var onGround pk.Boolean

	if err := p.Scan(&entityID, &x, &y, &z, &yaw, &pitch, &onGround); err != nil {
		return fmt.Errorf("unable to read EntityTeleport packet: %w", err)
	}

	return nil
}

func Advancements(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var action pk.Byte
	var data pk.String

	if err := p.Scan(&action, &data); err != nil {
		return fmt.Errorf("unable to read Advancements packet: %w", err)
	}

	return nil
}

func EntityProperties(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	/*var entityID pk.Int
	var count pk.VarInt
	var properties pk.Ary[]pk.String

	if err := p.Scan(&entityID, &properties); err != nil {
		return err
	}*/

	return nil
}

func EntityEffect(c *Client, p pk.Packet, cancel context.CancelFunc) error {
	var (
		entityID  pk.VarInt
		effectID  pk.VarInt
		amplifier pk.Byte
		duration  pk.VarInt
		flags     pk.Byte
	)

	if err := p.Scan(&entityID, &effectID, &amplifier, &duration, &flags); err != nil {
		return fmt.Errorf("unable to read EntityEffect packet: %w", err)
	}

	if _, ok := effects.ByID[int32(effectID)]; ok {
		c.Player.EntityPlayer.ActivePotionEffects[int32(effectID)] = effects.EffectStatus{
			ID:            int32(effectID),
			Amplifier:     byte(amplifier),
			Duration:      int32(duration),
			ShowParticles: flags&0x01 == 0x01,
			ShowIcon:      flags&0x04 == 0x04,
		}
	}
	return nil
}
