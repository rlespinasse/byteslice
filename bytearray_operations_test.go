package gobits

import (
	"reflect"
	"testing"
)

func TestRightShift(t *testing.T) {
	var val []byte
	tests := []struct {
		name   string
		data   []byte
		shift  uint64
		result []byte
	}{
		{"no shift to the right", []byte{0xDA, 0x99, 0xBA}, 0, []byte{0xDA, 0x99, 0xBA}},
		{"low shift to the right", []byte{0xDA, 0x99, 0xBA}, 1, []byte{0xB5, 0x33, 0x74}},
		{"middle shift to the right", []byte{0xDA, 0x99, 0xBA}, 8, []byte{0x99, 0xBA, 0x00}},
		{"high shift to the right", []byte{0xDA, 0x99, 0xBA}, 16, []byte{0xBA, 0x00, 0x00}},
		{"overflow shift to the right", []byte{0xDA, 0x99, 0xBA}, 24, []byte{0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {

		val = RightShift(tt.data, tt.shift)
		if !reflect.DeepEqual(val, tt.result) {
			t.Errorf("Test '%v' failed: RightShift(0x%x, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.shift,
				val,
				tt.result)
		}
	}
}

func BenchmarkRightShift(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RightShift([]byte{0x99, 0xBA}, 1)
	}
}

func BenchmarkRightShiftOver1Byte(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RightShift([]byte{0x99, 0xBA}, 9)
	}
}

func BenchmarkRightShiftOver1Byte(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RightShift([]byte{0x99, 0xBA}, 9)
	}
}

func TestLeftShift(t *testing.T) {
	var val []byte
	tests := []struct {
		name   string
		data   []byte
		shift  uint64
		result []byte
	}{
		{"no shift to the left", []byte{0xDA, 0x99, 0xBA}, 0, []byte{0xDA, 0x99, 0xBA}},
		{"low shift to the left", []byte{0xDA, 0x99, 0xBA}, 1, []byte{0x6D, 0x4C, 0xDD}},
		{"middle shift to the left", []byte{0xDA, 0x99, 0xBA}, 8, []byte{0x00, 0xDA, 0x99}},
		{"high shift to the left", []byte{0xDA, 0x99, 0xBA}, 16, []byte{0x00, 0x00, 0xDA}},
		{"overflow shift to the left", []byte{0xDA, 0x99, 0xBA}, 24, []byte{0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {

		val = LeftShift(tt.data, tt.shift)
		if !reflect.DeepEqual(val, tt.result) {
			t.Errorf("Test '%v' failed: LeftShift(0x%x, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.shift,
				val,
				tt.result)
		}
	}
}

func BenchmarkLeftShift(t *testing.B) {
	for i := 0; i < t.N; i++ {
		LeftShift([]byte{0x99, 0xBA}, 1)
	}
}

func BenchmarkLeftShiftOver1Byte(t *testing.B) {
	for i := 0; i < t.N; i++ {
		LeftShift([]byte{0x99, 0xBA}, 9)
	}
}

func TestExtractBytes(t *testing.T) {
	var val []byte
	tests := []struct {
		name                     string
		data                     []byte
		lsbPosition, msbPosition uint64
		result                   []byte
	}{
		{"extract nothing", []byte{0xDA, 0x99, 0xBA}, 0, 0, []byte{}},
		{"extract nothing due to inversed positions", []byte{0xDA, 0x99, 0xBA}, 16, 8, []byte{}},
		{"extract nothing due to wrong positions", []byte{0xDA, 0x99, 0xBA}, 100, 101, []byte{}},
		{"extract only in one byte", []byte{0xDA, 0x99, 0xBA}, 5, 7, []byte{0x05}},
		{"extract one byte over two bytes", []byte{0xDA, 0x99, 0xBA}, 7, 8, []byte{0x03}},
		{"extract two bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 6, 17, []byte{0x0A, 0x66}},
		{"extract three bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 1, 22, []byte{0x2D, 0x4C, 0xDD}},
		{"extract all bytes", []byte{0xDA, 0x99, 0xBA}, 0, 23, []byte{0xDA, 0x99, 0xBA}},
		{"extract all bytes with an overflow position", []byte{0xDA, 0x99, 0xBA}, 0, 100, []byte{0xDA, 0x99, 0xBA}},
	}
	for _, tt := range tests {

		val = ExtractBytes(tt.data, tt.lsbPosition, tt.msbPosition)
		if !reflect.DeepEqual(val, tt.result) {
			t.Errorf("Test '%v' failed: ExtractBytes(0x%x, %v, %v) was 0x%x, should be 0x%x",
				tt.name,
				tt.data, tt.lsbPosition, tt.msbPosition,
				val,
				tt.result)
		}
	}
}

func BenchmarkExtractBytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA}, 7, 8)
	}
}

func BenchmarkExtractBytesFrom3Bytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 15, 16)
	}
}

func BenchmarkExtract2ByteFrom3Bytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 8, 19)
	}
}

func BenchmarkExtractAllBytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 0, 23)
	}
}

func BenchmarkComputeSize(t *testing.B) {
	for i := 0; i < t.N; i++ {
		computeSize(7, 85)
	}
}

func BenchmarkTrim(t *testing.B) {
	for i := 0; i < t.N; i++ {
		trim([]byte{0x99, 0xBA, 0x00, 0x02}, 2)
	}
}
