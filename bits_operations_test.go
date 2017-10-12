package gobits

import (
	"testing"
)

func TestContainsBits(t *testing.T) {
	var val bool
	tests := []struct {
		name   string
		a, b   byte
		result bool
	}{
		{"low in hi", 0xf0, 5, false},
		{"hi in hi", 0xf0, 0x20, true},
		{"split", 0xf0, 0x22, false},
	}
	for _, tt := range tests {

		val = ContainsBits(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ContainsBits(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestSetBits(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		a, b   byte
		result byte
	}{
		{"set low bit", 0xf0, 0, 0xf0},
		{"high and low", 0xf0, 0x11, 0xf1},
	}
	for _, tt := range tests {

		val = SetBits(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: SetBits(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestUnsetBits(t *testing.T) {
	var val byte
	tests := []struct {
		name   string
		a, b   byte
		result byte
	}{
		{"unset zero bit", 0xf0, 7, 0xf0},
		{"unset some bits", 0xf0, 0x11, 0xe0},
	}
	for _, tt := range tests {

		val = UnsetBits(tt.a, tt.b)
		if val != tt.result {
			t.Errorf("Test '%v' failed: UnsetBits(%v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b,
				val,
				tt.result)
		}
	}
}

func TestExtractBits(t *testing.T) {
	var val byte
	tests := []struct {
		name    string
		a, b, c byte
		result  byte
	}{
		{"get bottom 3 bits", 0xf2, 0, 2, 2},
		{"get low bit", 0xf0, 0, 0, 0},
		{"lsb and msb out of order", 0xf0, 2, 0, 0},
	}
	for _, tt := range tests {

		val = ExtractBits(tt.a, tt.b, tt.c)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ExtractBits(%v, %v, %v) was %v, should be %v",
				tt.name,
				tt.a, tt.b, tt.c,
				val,
				tt.result)
		}
	}
}
