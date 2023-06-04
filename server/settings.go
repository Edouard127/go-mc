package server

import (
	"golang.org/x/time/rate"
	"time"
)

type ServerSettings struct {
	Name                        string `toml:"name"`
	MaxPlayers                  int    `toml:"max-players"`
	ViewDistance                int32  `toml:"view-distance"`
	ListenAddress               string `toml:"listen-address"`
	MessageOfTheDay             string `toml:"motd"`
	Icon                        string `toml:"icon"`
	NetworkCompressionThreshold int    `toml:"network-compression-threshold"`
	OnlineMode                  bool   `toml:"online-mode"`
	LevelName                   string `toml:"level-name"`
	EnforceSecureProfile        bool   `toml:"enforce-secure-profile"`

	ChunkLoadingLimiter       Limiter `toml:"chunk-loading-limiter"`
	PlayerChunkLoadingLimiter Limiter `toml:"player-chunk-loading-limiter"`
}

type Limiter struct {
	Every duration `toml:"every"`
	N     int
}

// Limiter convert this to *rate.Limiter
func (l *Limiter) Limiter() *rate.Limiter {
	return rate.NewLimiter(rate.Every(l.Every.Duration), l.N)
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return
}

var DefaultServerSettings = ServerSettings{
	Name:                        "Go-MC Server",
	MaxPlayers:                  20,
	ViewDistance:                10,
	ListenAddress:               "0.0.0.0:25565",
	MessageOfTheDay:             "Welcome to Go-MC Server",
	Icon:                        "",
	NetworkCompressionThreshold: 1,
	OnlineMode:                  true,
	LevelName:                   "world",
	EnforceSecureProfile:        false,
	ChunkLoadingLimiter: Limiter{
		Every: duration{Duration: 50 * time.Millisecond},
		N:     100,
	},
	PlayerChunkLoadingLimiter: Limiter{
		Every: duration{Duration: 50 * time.Millisecond},
		N:     100,
	},
}
