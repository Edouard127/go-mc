// Package provider bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at examples/ .
package provider

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	auth "github.com/maxsupermanhd/go-mc-ms-auth"
	"io"
	"net"
	"strconv"

	"github.com/Edouard127/go-mc/chat"
	"github.com/Edouard127/go-mc/data/packetid"
	mcnet "github.com/Edouard127/go-mc/net"
	pk "github.com/Edouard127/go-mc/net/packet"
)

// ProtocolVersion is the protocol version number of minecraft net protocol
const ProtocolVersion = 759
const DefaultPort = mcnet.DefaultPort

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err error) {
	return c.join(context.Background(), &mcnet.DefaultDialer, addr)
}

// JoinServerWithDialer is similar to JoinServer but using a Dialer.
func (c *Client) JoinServerWithDialer(d *net.Dialer, addr string) (err error) {
	dialer := (*mcnet.Dialer)(d)
	return c.join(context.Background(), dialer, addr)
}

func (c *Client) join(ctx context.Context, d *mcnet.Dialer, addr string) error {
	const Handshake = 0x00
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	var port int64

	if err != nil {
		_, records, err := net.LookupSRV("minecraft", "tcp", host)
		if err == nil && len(records) > 0 {
			addr = net.JoinHostPort(addr, strconv.Itoa(int(records[0].Port)))
			return c.join(ctx, d, addr)
		} else {
			addr = net.JoinHostPort(addr, strconv.Itoa(DefaultPort))
			return c.join(ctx, d, addr)
		}
	} else {
		port, err = strconv.ParseInt(portStr, 0, 16)
		if err != nil {
			return fmt.Errorf("parse port: %w", err)
		}
	}

	// Dial connection
	if c.Conn, err = d.DialMCContext(ctx, addr); err != nil {
		return fmt.Errorf("dial connection: %w", err)
	}

	// Handshake
	if err := c.Conn.WritePacket(pk.Marshal(
		Handshake,
		pk.VarInt(ProtocolVersion),
		pk.String(host),
		pk.UnsignedShort(port),
		pk.VarInt(2),
	)); err != nil {
		return fmt.Errorf("handshake: %w", err)
	}
	// Login Start
	if err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketLoginStart,
		pk.String(c.Auth.Name),
		pk.Opt{
			If: c.Auth.AsTk != "",
			Value: pk.Tuple{
				pk.Boolean(true),
				keyPair(c.Auth.KeyPair),
			},
			Else: pk.Boolean(false),
		},
	)); err != nil {
		return fmt.Errorf("login start: %w", err)
	}

	var p pk.Packet
	for {
		//Receive Packet
		c.Conn.ReadPacket(&p)

		//Handle Packet
		switch p.ID {
		case packetid.CPacketLoginDisconnect:
			var reason chat.Message
			if err := p.Scan(&reason); err != nil {
				break
			}
			return fmt.Errorf("login disconnect: %s", reason)

		case packetid.CPacketEncryptionRequest:
			if err := handleEncryptionRequest(c, p); err != nil {
				return fmt.Errorf("encryption request: %w", err)
			}

		case packetid.CPacketLoginSuccess:
			var (
				euuid pk.UUID
				name  pk.String
			)
			if err := p.Scan(&euuid, &name); err != nil {
				return fmt.Errorf("login success: %w", err)
			}
			c.Player.UUID = uuid.UUID(euuid)
			c.Player.Username = string(name)
			return nil
		case packetid.CPacketSetCompression:
			var threshold pk.VarInt
			if err := p.Scan(&threshold); err != nil {
				return fmt.Errorf("set compression: %w", err)
			}
			c.Conn.SetThreshold(int(threshold))
		case packetid.CPacketLoginPluginRequest:
			p.ID = packetid.SPacketLoginPluginResponse
			if err := c.Conn.WritePacket(p); err != nil {
				return fmt.Errorf("login plugin response: %w", err)
			}
		}
	}
}

type keyPair auth.KeyPair

func (k keyPair) WriteTo(w io.Writer) (int64, error) {
	block, _ := pem.Decode([]byte(k.KeyPair.PublicKey))
	if block == nil {
		return 0, errors.New("pem decode error: no data is found")
	}
	signature, err := base64.StdEncoding.DecodeString(k.PublicKeySignature)
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		pk.Long(k.ExpiresAt.UnixMilli()),
		pk.ByteArray(block.Bytes),
		pk.ByteArray(signature),
	}.WriteTo(w)
}
