package provider

import (
	"fmt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"time"
)

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (cl *Client) HandleGame() error {
	var p pk.Packet
	var err error
	start := time.Now()
	for {
		if err = cl.Conn.ReadPacket(&p); err != nil {
			panic(fmt.Errorf("read packet error 0x%x: %v", p.ID, err))
		}

		if err = cl.Events.HandlePacket(cl, p); err != nil {
			panic(fmt.Errorf("handle packet error 0x%x: %v", p.ID, err))
		}

		if time.Since(start) >= time.Millisecond*50 {
			if err = cl.Events.Tick(cl); err != nil {
				panic(fmt.Errorf("tick error: %v", err))
			}
			start = time.Now()
		}
	}
}
