package server

import (
	"github.com/Tnze/go-mc/bot/provider"
	"testing"
)

func TestExampleServer_Listen(t *testing.T) {
	server := NewServer(DefaultServerSettings)
	err := server.Listen("localhost:25565")
	if err != nil {
		t.Fatal(err)
	}
}

func TestExampleServer_Ping(t *testing.T) {
	go func() {
		server := NewServer(DefaultServerSettings)
		err := server.Listen("localhost:25565")
		if err != nil {
			t.Fatal(err)
		}
	}()
	data, _, err := provider.PingAndList("localhost:25565")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
