package provider

import (
	"fmt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"time"
)

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (cl *Client) HandleGame() error {
	var e1 error
	go func() {
		for {
			e1 = cl.Events.Tick(cl)
			if e1 != nil {
				panic(fmt.Errorf("tick error: %v", e1))
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()

	var p pk.Packet
	var e2 error
	for {
		if e2 = cl.Conn.ReadPacket(&p); e2 != nil {
			panic(fmt.Errorf("read packet error 0x%x: %v", p.ID, e2))
		}

		if e2 = cl.Events.HandlePacket(cl, p); e2 != nil {
			panic(fmt.Errorf("handle packet error %x: %v", p.ID, e2))
		}
	}
}
