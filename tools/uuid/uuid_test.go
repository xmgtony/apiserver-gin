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
	genUUID := GenUUID()
	uuid, err := ParseUUIDFromStr(genUUID)
	if err != nil || genUUID != uuid {
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
