package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Edouard127/go-mc/chat"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Edouard127/go-mc/data/packetid"
	mcnet "github.com/Edouard127/go-mc/net"
	pk "github.com/Edouard127/go-mc/net/packet"
)

type Status struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Description        chat.Message `json:"description"`
	Favicon            string       `json:"favicon"`
	EnforcesSecureChat bool         `json:"enforcesSecureChat"`
}

// PingAndList check server status and list online player.
// Returns a JSON data with server status, and the delay.
//
// For more information for JSON format, see https://wiki.vg/Server_List_Ping#Response
func PingAndList(addr string) (Status, time.Duration, error) {
	conn, err := mcnet.DialMC(addr)
	if err != nil {
		return Status{}, 0, fmt.Errorf("dial: %w", err)
	}
	return pingAndList(context.Background(), addr, conn)
}

// PingAndListTimeout is the version of PingAndList with max request time.
func PingAndListTimeout(addr string, timeout time.Duration) (Status, time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return PingAndListContext(ctx, addr)
}

func PingAndListContext(ctx context.Context, addr string) (data Status, delay time.Duration, err error) {
	var conn *mcnet.Conn
	conn, err = mcnet.DefaultDialer.DialMCContext(ctx, addr)
	if err != nil {
		return
	}
	return pingAndList(ctx, addr, conn)
}

func pingAndList(ctx context.Context, addr string, conn *mcnet.Conn) (data Status, delay time.Duration, err error) {
	if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
		if err = conn.Socket.SetDeadline(deadline); err != nil {
			return
		}
		defer func() {
			// Reset deadline
			if err2 := conn.Socket.SetDeadline(time.Time{}); err2 != nil {
				if err2 == nil {
					err = err2
				}
				return
			}
			// Map error type
			if errors.Is(err, os.ErrDeadlineExceeded) {
				err = context.DeadlineExceeded
			}
		}()
	}
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	var port int64

	if err != nil {
		_, records, err := net.LookupSRV("minecraft", "tcp", host)
		if err == nil && len(records) > 0 {
			addr = net.JoinHostPort(addr, strconv.Itoa(int(records[0].Port)))
			return pingAndList(ctx, addr, conn)
		} else {
			addr = net.JoinHostPort(addr, strconv.Itoa(DefaultPort))
			return pingAndList(ctx, addr, conn)
		}
	} else {
		port, err = strconv.ParseInt(portStr, 0, 16)
		if err != nil {
			return
		}
	}

	const Handshake = 0x00
	//握手
	if err := conn.WritePacket(pk.Marshal(
		Handshake,                  //Handshake packet ID
		pk.VarInt(ProtocolVersion), //Protocol version
		pk.String(host),            //Server's address
		pk.UnsignedShort(port),
		pk.Byte(1),
	)); err != nil {
		return Status{}, 0, fmt.Errorf("bot: send handshake packet fail: %v", err)
	}

	//LIST
	//请求服务器状态
	if err := conn.WritePacket(pk.Marshal(
		packetid.CPacketEncryptionRequest,
	)); err != nil {
		return Status{}, 0, fmt.Errorf("bot: send list packet fail: %v", err)
	}

	var p pk.Packet
	//服务器返回状态
	if err := conn.ReadPacket(&p); err != nil {
		return Status{}, 0, fmt.Errorf("bot: recv list packect fail: %v", err)
	}
	var s pk.String
	err = p.Scan(&s)
	if err != nil {
		return Status{}, 0, fmt.Errorf("bot: scan list packet fail: %v", err)
	}

	//PING
	startTime := time.Now()
	if err := conn.WritePacket(pk.Marshal(
		packetid.SPacketPingRequest,
		pk.Long(startTime.Unix()),
	)); err != nil {
		return Status{}, 0, fmt.Errorf("bot: send ping packet fail: %v", err)
	}

	if err := conn.ReadPacket(&p); err != nil {
		return Status{}, 0, fmt.Errorf("bot: recv pong packet fail: %v", err)
	}
	var t pk.Long
	err = p.Scan(&t)
	if err != nil {
		return Status{}, 0, fmt.Errorf("bot: scan pong packet fail: %v", err)
	}
	if t != pk.Long(startTime.Unix()) {
		return Status{}, 0, fmt.Errorf("bot: pong packet no match: %v", err)
	}

	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		return
	}

	return data, time.Since(startTime), nil
}
