package gobits

import "testing"

func TestContainsBits(t *testing.T) {
	var val bool
	tests := []struct {
		name       string
		data, bits byte
		result     bool
	}{
		{"low in hi", 0xf0, 0x05, false},
		{"hi in hi", 0xf0, 0x20, true},
		{"split", 0xf0, 0x22, false},
	}
	for _, tt := range tests {

		val = ContainsBits(tt.data, tt.bits)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ContainsBits(0x%x, 0x%x) was %v, should be %v",
				tt.name,
				tt.data, tt.bits,
				val,
				tt.result)
		}
	}
}

func BenchmarkContainsBitsTrue(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ContainsBits(0xf0, 0x20)
	}
}

func BenchmarkContainsBitsFalse(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ContainsBits(0xf0, 5)
	}
}

func BenchmarkContainsBitsFalseSplit(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ContainsBits(0xf0, 0x22)
	}
}

func TestSetBits(t *testing.T) {
	var val byte
	tests := []struct {
		name       string
		data, bits byte
		result     byte
	}{
		{"set low bit", 0xf0, 0x00, 0xf0},
		{"high and low", 0xf0, 0x11, 0xf1},
	}
	for _, tt := range tests {

		val = SetBits(tt.data, tt.bits)
		if val != tt.result {
			t.Errorf("Test '%v' failed: SetBits(0x%x, 0x%x) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.bits,
				val,
				tt.result)
		}
	}
}

func BenchmarkSetBitsLow(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SetBits(0xf0, 0)
	}
}

func BenchmarkSetBitsHighLow(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SetBits(0xf0, 0x11)
	}
}

func TestUnsetBits(t *testing.T) {
	var val byte
	tests := []struct {
		name       string
		data, bits byte
		result     byte
	}{
		{"unset zero bit", 0xf0, 0x07, 0xf0},
		{"unset some bits", 0xf0, 0x11, 0xe0},
	}
	for _, tt := range tests {

		val = UnsetBits(tt.data, tt.bits)
		if val != tt.result {
			t.Errorf("Test '%v' failed: UnsetBits(0x%x, 0x%x) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.bits,
				val,
				tt.result)
		}
	}
}

func BenchmarkUnsetBitsZeroBits(t *testing.B) {
	for i := 0; i < t.N; i++ {
		UnsetBits(0xf0, 7)
	}
}

func BenchmarkUnsetBitsSomeBits(t *testing.B) {
	for i := 0; i < t.N; i++ {
		UnsetBits(0xf0, 0x11)
	}
}

func TestExtractBits(t *testing.T) {
	var val byte
	tests := []struct {
		name                     string
		data                     byte
		lsbPosition, msbPosition uint8
		result                   byte
	}{
		{"get bottom 3 bits", 0xf2, 0, 2, 0x02},
		{"get low bit", 0xf0, 0, 0, 0x00},
		{"lsb and msb out of order", 0xf0, 2, 0, 0x00},
	}
	for _, tt := range tests {

		val = ExtractBits(tt.data, tt.lsbPosition, tt.msbPosition)
		if val != tt.result {
			t.Errorf("Test '%v' failed: ExtractBits(0x%x, %v, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.lsbPosition, tt.msbPosition,
				val,
				tt.result)
		}
	}
}

func BenchmarkExtractBitsBottom(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBits(0xf0, 0, 2)
	}
}

func BenchmarkExtractBitsLow(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBits(0xf0, 0, 0)
	}
}

func BenchmarkExtractBitsOutOfOrder(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBits(0xf0, 2, 0)
	}
}
