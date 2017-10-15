package gobits

import (
	"testing"
)

var testcasesContainsBit = []struct {
	name   string
	data   byte
	bit    uint8
	result bool
}{
	{"hi in hi", 0xf0, 5, true},
	{"lo in hi", 0xf0, 2, false},
}

func TestContainsBit(t *testing.T) {
	var val bool
	for _, tc := range testcasesContainsBit {
		t.Run(tc.name, func(t *testing.T) {
			val = ContainsBit(tc.data, tc.bit)
			if val != tc.result {
				t.Errorf("ContainsBit(%x, %v) was %v, should be %v",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkContainsBit(b *testing.B) {
	var val bool
	for _, tc := range testcasesContainsBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ContainsBit(tc.data, tc.bit)
			}
		})
	}
}

var testcasesSetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"set low bit", 0xf0, 0, 0xf1},
	{"set already-set bit", 0xf0, 5, 0xf0},
}

func TestSetBit(t *testing.T) {
	var val byte
	for _, tc := range testcasesSetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = SetBit(tc.data, tc.bit)
			if val != tc.result {
				t.Errorf("SetBit(%x, %v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkSetBit(b *testing.B) {
	var val byte
	for _, tc := range testcasesSetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = SetBit(tc.data, tc.bit)
			}
		})
	}
}

var testcasesUnsetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"unset high bit", 0xf0, 7, 0x70},
	{"unset already-clear bit", 0xf0, 0, 0xf0},
}

func TestUnsetBit(t *testing.T) {
	var val byte
	for _, tc := range testcasesUnsetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = UnsetBit(tc.data, tc.bit)
			if val != tc.result {
				t.Errorf("UnsetBit(%x, %v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkUnsetBit(b *testing.B) {
	var val byte
	for _, tc := range testcasesUnsetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = UnsetBit(tc.data, tc.bit)
			}
		})
	}
}

var testcasesGetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"get low bit of high nibble", 0xf0, 4, 0x10},
	{"get low bit", 0xf0, 0, 0},
}

func TestGetBit(t *testing.T) {
	var val byte
	for _, tc := range testcasesGetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = GetBit(tc.data, tc.bit)
			if val != tc.result {
				t.Errorf("GetBit(%x, %v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkGetBit(b *testing.B) {
	var val byte
	for _, tc := range testcasesGetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = GetBit(tc.data, tc.bit)
			}
		})
	}
}
