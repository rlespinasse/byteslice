package byteslice

import (
	"testing"
)

var testcasesByteItemContainsBit = []struct {
	name   string
	data   byte
	bit    uint8
	result bool
}{
	{"hi in hi", 0xf0, 5, true},
	{"lo in hi", 0xf0, 2, false},
}

func TestByteItemContainsBit(t *testing.T) {
	var val bool
	for _, tc := range testcasesByteItemContainsBit {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).ContainsBit(tc.bit)
			if val != tc.result {
				t.Errorf("ByteItem(%x).ContainsBit(%v) was %v, should be %v",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemContainsBit(b *testing.B) {
	var val bool
	for _, tc := range testcasesByteItemContainsBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).ContainsBit(tc.bit)
			}
		})
	}
}

var testcasesByteItemContainsBits = []struct {
	name       string
	data, bits byte
	result     bool
}{
	{"low in hi", 0xf0, 0x05, false},
	{"hi in hi", 0xf0, 0x20, true},
	{"split", 0xf0, 0x22, false},
}

func TestByteItemContains(t *testing.T) {
	var val bool
	for _, tc := range testcasesByteItemContainsBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).Contains(ByteItem(tc.bits))
			if val != tc.result {
				t.Errorf("ByteItem(%x).Contains(%x) was %v, should be %v",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemContains(b *testing.B) {
	var val bool
	for _, tc := range testcasesByteItemContainsBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).Contains(ByteItem(tc.bits))
			}
		})
	}
}

var testcasesByteItemSetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"set low bit", 0xf0, 0, 0xf1},
	{"set already-set bit", 0xf0, 5, 0xf0},
}

func TestByteItemSetBit(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemSetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).SetBit(tc.bit)
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).SetBit(%v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemSetBit(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemSetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).SetBit(tc.bit)
			}
		})
	}
}

var testcasesByteItemSetBits = []struct {
	name       string
	data, bits byte
	result     byte
}{
	{"set low bit", 0xf0, 0x00, 0xf0},
	{"high and low", 0xf0, 0x11, 0xf1},
}

func TestByteItemSet(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemSetBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).Set(ByteItem(tc.bits))
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).Set(%x) was %x, should be %x",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemSet(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemSetBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).Set(ByteItem(tc.bits))
			}
		})
	}
}

var testcasesByteItemUnsetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"unset high bit", 0xf0, 7, 0x70},
	{"unset already-clear bit", 0xf0, 0, 0xf0},
}

func TestByteItemUnsetBit(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemUnsetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).UnsetBit(tc.bit)
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).UnsetBit(%v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemUnsetBit(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemUnsetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).UnsetBit(tc.bit)
			}
		})
	}
}

var testcasesByteItemUnsetBits = []struct {
	name       string
	data, bits byte
	result     byte
}{
	{"unset zero bit", 0xf0, 0x07, 0xf0},
	{"unset some bits", 0xf0, 0x11, 0xe0},
}

func TestByteItemUnset(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemUnsetBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).Unset(ByteItem(tc.bits))
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).Unset(%x) was %x, should be %x",
					tc.data, tc.bits,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemUnset(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemUnsetBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).Unset(ByteItem(tc.bits))
			}
		})
	}
}

var testcasesByteItemGetBit = []struct {
	name   string
	data   byte
	bit    uint8
	result byte
}{
	{"get low bit of high nibble", 0xf0, 4, 0x10},
	{"get low bit", 0xf0, 0, 0},
}

func TestByteItemGetBit(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemGetBit {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).GetBit(tc.bit)
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).GetBit(%v) was %x, should be %x",
					tc.data, tc.bit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemGetBit(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemGetBit {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).GetBit(tc.bit)
			}
		})
	}
}

var testcasesByteItemSubsetBits = []struct {
	name                     string
	data                     byte
	lsbPosition, msbPosition uint8
	result                   byte
}{
	{"get bottom 3 bits", 0xf2, 0, 2, 0x02},
	{"get low bit", 0xf0, 0, 0, 0x00},
	{"lsb and msb out of order", 0xf0, 2, 0, 0x00},
}

func TestByteItemSubset(t *testing.T) {
	var val ByteItem
	for _, tc := range testcasesByteItemSubsetBits {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteItem(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			if val != ByteItem(tc.result) {
				t.Errorf("ByteItem(%x).Subset(%v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteItemSubset(b *testing.B) {
	var val ByteItem
	for _, tc := range testcasesByteItemSubsetBits {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteItem(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
