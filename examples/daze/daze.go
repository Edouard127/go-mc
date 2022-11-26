// Daze could join an offline-mode server as client.
// Just standing there and do nothing. Automatically reborn after five seconds of death.
package main

import (
	"errors"
	"flag"
	"log"
	"time"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/item"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
	"github.com/Tnze/go-mc/level"
)

var address = flag.String("address", "127.0.0.1", "The server address")
var name = flag.String("name", "Daze", "The player's name")
var playerID = flag.String("uuid", "", "The player's UUID")
var accessToken = flag.String("token", "", "AccessToken")

var client *bot.Client
var player *bot.Player
var worldManager *bot.World
var screenManager *bot.Manager

func main() {
	flag.Parse()
	//log.SetOutput(colorable.NewColorableStdout())
	client = bot.NewClient()
	client.Auth = bot.Auth{
		Name: *name,
		UUID: *playerID,
		AsTk: *accessToken,
	}
	player = bot.NewPlayer(client, bot.DefaultSettings)
	bot.EventsListener{
		c:            client,
		GameStart:    onGameStart,
		ChatMsg:      onChatMsg,
		SystemMsg:    onSystemMsg,
		Disconnect:   onDisconnect,
		HealthChange: onHealthChange,
		Death:        onDeath,
	}.Attach(client)
	worldManager = bot.NewWorld(client, player, world.EventsListener{
		LoadChunk:   onChunkLoad,
		UnloadChunk: onChunkUnload,
	})
	screenManager = bot.NewManager(client, screen.EventsListener{
		Open:    nil,
		SetSlot: onScreenSlotChange,
		Close:   nil,
	})

	//Login
	err := client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//JoinGame
	for {
		if err = client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}

		if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
			if err := new(DisconnectErr); errors.As(err2, err) {
				log.Print("Disconnect: ", err.Reason)
				return
			} else {
				// print and ignore the error
				log.Print(err2)
			}
		} else {
			log.Fatal(err)
		}
	}
}

func onDeath() error {
	log.Println("Died and Respawned")
	// If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	go func() {
		time.Sleep(time.Second * 5)
		err := player.Respawn()
		if err != nil {
			log.Print(err)
		}
	}()
	return nil
}

func onGameStart() error {
	log.Println("Game start")
	return nil //if err isn't nil, HandleGame() will return it.
}

func onChatMsg(c *bot.PlayerMessage) error {
	log.Println("Chat:", c.SignedMessage.String())
	return nil
}

func onSystemMsg(c chat.Message, pos byte) error {
	log.Printf("System: %v, Location: %v", c, pos)
	return nil
}

func onChunkLoad(pos level.ChunkPos) error {
	log.Println("Load chunk:", pos)
	return nil
}

func onChunkUnload(pos level.ChunkPos) error {
	log.Println("Unload chunk:", pos)
	return nil
}

func onScreenSlotChange(id, index int) error {
	if id == -2 {
		log.Printf("Slot: inventory: %v", screenManager.Inventory.Slots[index])
	} else if id == -1 && index == -1 {
		log.Printf("Slot: cursor: %v", screenManager.Cursor)
	} else {
		container, ok := screenManager.Screens[id]
		if ok {
			// Currently, only inventory container is supported
			switch container.(type) {
			case *bot.Inventory:
				slot := container.(*bot.Inventory).Slots[index]
				itemInfo := item.ByID[item.ID(slot.ID)]
				log.Printf("Slot: Screen[%d].Slot[%d]: [%v] * %d | NBT: %v", id, index, itemInfo, slot.Count, slot.NBT)
			}
		}
	}
	return nil
}

func onHealthChange(health float32) error {
	log.Printf("HealthChange: %v", health)
	return nil
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}

func onDisconnect(reason chat.Message) error {
	// return an error value so that we can stop main loop
	return DisconnectErr{Reason: reason}
}
