package gobits

import (
	"testing"
)

func TestContainsBit(t *testing.T) {
	var val bool
	tests := []struct {
		name   string
		a, b   byte
		result bool
	}{
		{"hi in hi", 0xf0, 5, true},
		{"lo in hi", 0xf0, 2, false},
	}
	for _, tt := range tests {

		val = ContainsBit(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ContainsBit(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestSetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		a, b   byte
		result byte
	}{
		{"set low bit", 0xf0, 0, 0xf1},
		{"set already-set bit", 0xf0, 5, 0xf0},
	}
	for _, tt := range tests {

		val = SetBit(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: SetBit(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestUnsetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		a, b   byte
		result byte
	}{
		{"unset high bit", 0xf0, 7, 0x70},
		{"unset already-clear bit", 0xf0, 0, 0xf0},
	}
	for _, tt := range tests {

		val = UnsetBit(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: UnsetBit(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestGetBit(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		a, b   byte
		result byte
	}{
		{"get low bit of high nibble", 0xf0, 4, 0x10},
		{"get low bit", 0xf0, 0, 0},
	}
	for _, tt := range tests {

		val = GetBit(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: GetBit(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}
