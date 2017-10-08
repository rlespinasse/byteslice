package gobits

import (
	"testing"
)

func TestContainsBits(t *testing.T) {
	var result = ContainsBits(0xF0, 0xA0)
	if !result {
		t.Errorf("ContainsBits don't work")
	}
}

func TestSetBits(t *testing.T) {
	var result = SetBits(0xF0, 0x01)
	if result != 0xF1 {
		t.Errorf("SetBits don't work: %08b", result)
	}
}

func TestUnsetBits(t *testing.T) {
	var result = UnsetBits(0xF1, 0x01)
	if result != 0xF0 {
		t.Errorf("UnsetBits don't work: %08b", result)
	}
}
