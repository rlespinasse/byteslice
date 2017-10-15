package gobits

import "testing"

var testcasesContainsBits = []struct {
	name       string
	data, bits byte
	result     bool
}{
	{"low in hi", 0xf0, 0x05, false},
	{"hi in hi", 0xf0, 0x20, true},
	{"split", 0xf0, 0x22, false},
}

func TestContainsBits(t *testing.T) {
	var val bool
	for _, tc := range testcasesContainsBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ContainsBits(tc.data, tc.bits)
			if val != tc.result {
				t.Errorf("ContainsBits(%x, %x) was %v, should be %v",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkContainsBits(b *testing.B) {
	var val bool
	for _, tc := range testcasesContainsBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ContainsBits(tc.data, tc.bits)
			}
		})
	}
}

var testcasesSetBits = []struct {
	name       string
	data, bits byte
	result     byte
}{
	{"set low bit", 0xf0, 0x00, 0xf0},
	{"high and low", 0xf0, 0x11, 0xf1},
}

func TestSetBits(t *testing.T) {
	var val byte
	for _, tc := range testcasesSetBits {
		t.Run(tc.name, func(t *testing.T) {
			val = SetBits(tc.data, tc.bits)
			if val != tc.result {
				t.Errorf("SetBits(%x, %x) was %x, should be %x",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkSetBits(b *testing.B) {
	var val byte
	for _, tc := range testcasesSetBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = SetBits(tc.data, tc.bits)
			}
		})
	}
}

var testcasesUnsetBits = []struct {
	name       string
	data, bits byte
	result     byte
}{
	{"unset zero bit", 0xf0, 0x07, 0xf0},
	{"unset some bits", 0xf0, 0x11, 0xe0},
}

func TestUnsetBits(t *testing.T) {
	var val byte
	for _, tc := range testcasesUnsetBits {
		t.Run(tc.name, func(t *testing.T) {
			val = UnsetBits(tc.data, tc.bits)
			if val != tc.result {
				t.Errorf("UnsetBits(%x, %x) was %x, should be %x",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkUnsetBits(b *testing.B) {
	var val byte
	for _, tc := range testcasesUnsetBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = UnsetBits(tc.data, tc.bits)
			}
		})
	}
}

var testcasesExtractBits = []struct {
	name                     string
	data                     byte
	lsbPosition, msbPosition uint8
	result                   byte
}{
	{"get bottom 3 bits", 0xf2, 0, 2, 0x02},
	{"get low bit", 0xf0, 0, 0, 0x00},
	{"lsb and msb out of order", 0xf0, 2, 0, 0x00},
}

func TestExtractBits(t *testing.T) {
	var val byte
	for _, tc := range testcasesExtractBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ExtractBits(tc.data, tc.lsbPosition, tc.msbPosition)
			if val != tc.result {
				t.Errorf("ExtractBits(%x, %x) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkExtractBits(b *testing.B) {
	var val byte
	for _, tc := range testcasesExtractBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ExtractBits(tc.data, tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
