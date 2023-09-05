package realms

import (
	"fmt"
	"github.com/Edouard127/go-mc/auth/microsoft"
	"testing"
)

const Version = "1.20.1"

func TestNewRealms(t *testing.T) {
	account := microsoft.LoginFromCache(nil)

	realms := NewRealms(Version, account.Name, account.AccessToken, account.UUID)

	fmt.Printf("Realms available: %t, compatible: %s\n", unwrap(realms.Available()), unwrap(realms.Compatible()))

	for _, v := range unwrap(realms.Worlds()) {
		fmt.Printf("World %s: %d\n", v.Name, v.ID)
	}
}

func unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func noerr(err error) {
	if err != nil {
		panic(err)
	}
}
