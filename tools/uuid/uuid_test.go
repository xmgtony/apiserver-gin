package uuid

import (
	"testing"
)

func TestGenUUID(t *testing.T) {
	uuid := GenUUID()
	if len(uuid) <= 0 {
		t.Fatal("failed")
	}
	t.Log(uuid)
}

func TestGenUUIDFromStr(t *testing.T) {
	uuid, err := ParseUUIDFromStr(GenUUID())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuid)
}

func TestGenUUID16(t *testing.T) {
	uuid := GenUUID16()
	if len(uuid) <= 0 {
		t.Fatal("failed")
	}
	t.Log(uuid)
}
