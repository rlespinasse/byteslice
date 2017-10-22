package gobits

import (
	"reflect"
	"testing"
)

var testcasesRightShift = []struct {
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

func TestRightShift(t *testing.T) {
	var val []byte
	for _, tc := range testcasesRightShift {
		t.Run(tc.name, func(t *testing.T) {
			val = RightShift(tc.data, tc.shift)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RightShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRightShift(b *testing.B) {
	var val []byte
	for _, tc := range testcasesRightShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RightShift(tc.data, tc.shift)
			}
		})
	}
}

var testcasesLeftShift = []struct {
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

func TestLeftShift(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLeftShift {
		t.Run(tc.name, func(t *testing.T) {
			val = LeftShift(tc.data, tc.shift)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LeftShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

var testcasesMask = []struct {
	name   string
	data   []byte
	mask   []byte
	result []byte
}{
  {"test", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xaa}},
	{"longer mask than data", []byte{0xDA, 0x99, 0xBA},
		[]byte{0xAD, 0x11, 0xAB, 0x88, 0x11, 0xaa}, []byte{0x88, 0x11, 0xaa}},
	{"empty array", []byte{}, []byte{}, []byte{}},
}

func TestMask(t *testing.T) {
	var val []byte
	for _, tc := range testcasesMask {
		t.Run(tc.name, func(t *testing.T) {
			val = Mask(tc.data, tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("Mask(%x, %x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

var testcasesInclusiveMerge = []struct {
	name   string
	data   []byte
	mask   []byte
	result []byte
}{
	{"test", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xff, 0x99, 0xbb}},
	{"longer mask than data", []byte{0xDA, 0x99, 0xBA},
		[]byte{0xAD, 0x11, 0xAB, 0xff, 0x99, 0xbb}, []byte{0xff, 0x99, 0xbb}},
	{"empty array", []byte{}, []byte{}, []byte{}},
}

func TestInclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesInclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = InclusiveMerge(tc.data, tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("InclusiveMerge(%x, %x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

var testcasesExclusiveMerge = []struct {
	name   string
	data   []byte
	mask   []byte
	result []byte
}{
	{"test", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"longer mask than data", []byte{0xDA, 0x99, 0xBA},
		[]byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0x77, 0x88, 0x11}},
	{"empty array", []byte{}, []byte{}, []byte{}},
}

func TestExclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesExclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = ExclusiveMerge(tc.data, tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("ExclusiveMerge(%x, %x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

var testcasesNot = []struct {
	name   string
	data   []byte
	result []byte
}{
	{"test", []byte{0xDA, 0x99, 0xBA}, []byte{0x25, 0x66, 0x45}},
}

func TestNot(t *testing.T) {
	var val []byte
	for _, tc := range testcasesNot {
		t.Run(tc.name, func(t *testing.T) {
			val = Not(tc.data)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("Not(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}


func BenchmarkLeftShift(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLeftShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LeftShift(tc.data, tc.shift)
			}
		})
	}
}

var testcasesExtractBytes = []struct {
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

func TestExtractBytes(t *testing.T) {
	var val []byte
	for _, tc := range testcasesExtractBytes {
		t.Run(tc.name, func(t *testing.T) {
			val = ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("ExtractBytes(%x, %v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkExtractBytes(b *testing.B) {
	var val []byte
	for _, tc := range testcasesExtractBytes {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
