package gobits

import (
	"testing"
)

func TestContainsBit(t *testing.T) {
	var result = ContainsBit(0xF0, 5)
	if !result {
		t.Errorf("ContainsBit don't work")
	}
}

func TestSetBit(t *testing.T) {
	var result = SetBit(0xF0, 0)
	if result != 0xF1 {
		t.Errorf("SetBit don't work: %08b", result)
	}
}

func TestUnsetBit(t *testing.T) {
	var result = UnsetBit(0xF1, 0)
	if result != 0xF0 {
		t.Errorf("UnsetBit don't work: %08b", result)
	}
}

func TestGetBit(t *testing.T) {
	var result = GetBit(0xF0, 4)
	if result != 0x10 {
		t.Errorf("GetBit don't work: %08b", result)
	}
}
