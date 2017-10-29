package byteslice

import (
	"testing"
)

var tcRBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"get low bit of high nibble", 0xf0, 4, 0x10},
	{"get low bit", 0xf0, 0, 0},
}

func TestRBit(t *testing.T) {
	var val byte
	for _, tc := range tcRBit {
		t.Run(tc.name, func(t *testing.T) {
			val = RBit(tc.data, tc.bit)
			if val != (tc.result) {
				t.Errorf("RBit(%x, %v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRBit(b *testing.B) {
	var val byte
	for _, tc := range tcRBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RBit(tc.data, tc.bit)
			}
		})
	}
}

var tcRBitsSubset = []struct {
	name                                    string
	data                                    byte
	leastSignificantBit, mostSignificantBit uint8
	result                                  byte
}{
	{"get bottom 3 bits", 0xf2, 0, 2, 0x02},
	{"get low bit", 0xf0, 0, 0, 0x00},
	{"lsb and msb out of order", 0xf0, 2, 0, 0x00},
}

func TestRBitsSubset(t *testing.T) {
	var val byte
	for _, tc := range tcRBitsSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = RBitsSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			if val != tc.result {
				t.Errorf("RBitsSubset(%x, %v, %v) was %x, should be %x",
					tc.data, tc.leastSignificantBit, tc.mostSignificantBit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRBitsSubset(b *testing.B) {
	var val byte
	for _, tc := range tcRBitsSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RBitsSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			}
		})
	}
}
