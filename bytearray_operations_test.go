package gobits

import (
	"reflect"
	"testing"
)

func TestRightShift(t *testing.T) {
	var result = RightShift([]byte{0x99, 0xBA}, 1)
	if !reflect.DeepEqual(result, []byte{0x33, 0x74}) {
		t.Errorf("RightShift don't work: %x", result)
	}
}

func BenchmarkRightShift(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RightShift([]byte{0x99, 0xBA}, 1)
	}
}

func TestRightShiftOver1Byte(t *testing.T) {
	var result = RightShift([]byte{0x99, 0xBA}, 9)
	if !reflect.DeepEqual(result, []byte{0x74, 0x00}) {
		t.Errorf("RightShiftOver1Byte don't work: %x", result)
	}
}

func BenchmarkRightShiftOver1Byte(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RightShift([]byte{0x99, 0xBA}, 9)
	}
}

func TestLeftShift(t *testing.T) {
	var result = LeftShift([]byte{0x99, 0xBA}, 1)
	if !reflect.DeepEqual(result, []byte{0x4C, 0xDD}) {
		t.Errorf("LeftShift don't work: %x", result)
	}
}

func BenchmarkLeftShift(t *testing.B) {
	for i := 0; i < t.N; i++ {
		LeftShift([]byte{0x99, 0xBA}, 1)
	}
}

func TestLeftShiftOver1Byte(t *testing.T) {
	var result = LeftShift([]byte{0x99, 0xBA}, 9)
	if !reflect.DeepEqual(result, []byte{0x00, 0x4C}) {
		t.Errorf("LeftShiftOver1Byte don't work: %x", result)
	}
}

func BenchmarkLeftShiftOver1Byte(t *testing.B) {
	for i := 0; i < t.N; i++ {
		LeftShift([]byte{0x99, 0xBA}, 9)
	}
}

func TestExtract1ByteFrom2Bytes(t *testing.T) {
	var result = ExtractBytes([]byte{0x99, 0xBA}, 7, 8)
	if !reflect.DeepEqual(result, []byte{0x03}) {
		t.Errorf("Extract1ByteFrom2Bytes don't work: %x", result)
	}
}

func BenchmarkExtractBytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA}, 7, 8)
	}
}

func TestExtract1ByteFrom3Bytes(t *testing.T) {
	var result = ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 15, 16)
	if !reflect.DeepEqual(result, []byte{0x03}) {
		t.Errorf("Extract1ByteFrom3Bytes don't work: %x", result)
	}
}

func BenchmarkExtractBytesFrom3Bytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 15, 16)
	}
}

func TestExtract2ByteFrom3Bytes(t *testing.T) {
	var result = ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 8, 19)
	if !reflect.DeepEqual(result, []byte{0x09, 0xBA}) {
		t.Errorf("Extract2ByteFrom3Bytes don't work: %x", result)
	}
}

func BenchmarkExtract2ByteFrom3Bytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 8, 19)
	}
}

func TestExtractAllBytes(t *testing.T) {
	var result = ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 0, 23)
	if !reflect.DeepEqual(result, []byte{0x99, 0xBA, 0xDE}) {
		t.Errorf("ExtractAllBytes don't work: %x", result)
	}
}

func BenchmarkExtractAllBytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ExtractBytes([]byte{0x99, 0xBA, 0xDE}, 0, 23)
	}
}

func TestComputeSize(t *testing.T) {
	var result = ComputeSize(8, 75)
	if result != 9 {
		t.Errorf("ComputeSize don't work: %d", result)
	}
}

func BenchmarkComputeSize(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ComputeSize(8, 75)
	}
}

func TestTrim(t *testing.T) {
	var result = Trim([]byte{0x99, 0xBA, 0x00, 0x02}, 2)
	if !reflect.DeepEqual(result, []byte{0x00, 0x02}) {
		t.Errorf("TestTrim don't work: %x", result)
	}
}

func BenchmarkTrim(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Trim([]byte{0x99, 0xBA, 0x00, 0x02}, 2)
	}
}
