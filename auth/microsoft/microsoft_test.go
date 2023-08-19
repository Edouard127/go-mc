package microsoft

import (
	"testing"
)

func TestLoginDevice(t *testing.T) {
	xbox, err := LoginWithDeviceCode()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(xbox)

	mc, err := MinecraftLogin(xbox, false)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mc)
}
