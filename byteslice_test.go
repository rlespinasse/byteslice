package byteslice

import (
	"reflect"
	"testing"
)

func TestByteSliceReverse(t *testing.T) {
	data := []byte{0x01, 0x10}
	result := []byte{0x10, 0x01}

	val := ByteSlice(data).Reverse()
	if !reflect.DeepEqual(val, ByteSlice(result)) {
		t.Errorf("ByteSlice(%x).Reverse() was %x, should be %x",
			data,
			val,
			result)
	}
}

var testcasesByteSliceLeftShift = []struct {
	name   string
	data   []byte
	shift  uint64
	result []byte
}{
	{"no shift to the left", []byte{0xDA, 0x99, 0xBA}, 0, []byte{0xDA, 0x99, 0xBA}},
	{"low shift to the left", []byte{0xDA, 0x99, 0xBA}, 1, []byte{0xB5, 0x33, 0x74}},
	{"middle shift to the left", []byte{0xDA, 0x99, 0xBA}, 8, []byte{0x99, 0xBA, 0x00}},
	{"high shift to the left", []byte{0xDA, 0x99, 0xBA}, 16, []byte{0xBA, 0x00, 0x00}},
	{"overflow shift to the left", []byte{0xDA, 0x99, 0xBA}, 24, []byte{0x00, 0x00, 0x00}},
}

func TestByteSliceLeftShift(t *testing.T) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceLeftShift {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteSlice(tc.data).LeftShift(tc.shift)
			if !reflect.DeepEqual(val, ByteSlice(tc.result)) {
				t.Errorf("LeftShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteSliceLeftShift(b *testing.B) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceLeftShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteSlice(tc.data).LeftShift(tc.shift)
			}
		})
	}
}

var testcasesByteSliceRightShift = []struct {
	name   string
	data   []byte
	shift  uint64
	result []byte
}{
	{"no shift to the right", []byte{0xDA, 0x99, 0xBA}, 0, []byte{0xDA, 0x99, 0xBA}},
	{"low shift to the right", []byte{0xDA, 0x99, 0xBA}, 1, []byte{0x6D, 0x4C, 0xDD}},
	{"middle shift to the right", []byte{0xDA, 0x99, 0xBA}, 8, []byte{0x00, 0xDA, 0x99}},
	{"high shift to the right", []byte{0xDA, 0x99, 0xBA}, 16, []byte{0x00, 0x00, 0xDA}},
	{"overflow shift to the right", []byte{0xDA, 0x99, 0xBA}, 24, []byte{0x00, 0x00, 0x00}},
}

func TestByteSliceRightShift(t *testing.T) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceRightShift {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteSlice(tc.data).RightShift(tc.shift)
			if !reflect.DeepEqual(val, ByteSlice(tc.result)) {
				t.Errorf("RightShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteSliceRightShift(b *testing.B) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceRightShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteSlice(tc.data).RightShift(tc.shift)
			}
		})
	}
}

func TestByteSliceMaskError(t *testing.T) {
	val, err := ByteSlice([]byte{0x00, 0x00}).Mask(ByteSlice([]byte{0x00}))
	if err == nil || val != nil {
		t.Errorf("Mask with two byte arrays of different size needs to return an error and no value")
	}
}

func TestByteSliceInclusiveMergeError(t *testing.T) {
	val, err := ByteSlice([]byte{0x00, 0x00}).InclusiveMerge(ByteSlice([]byte{0x00}))
	if err == nil || val != nil {
		t.Errorf("InclusiveMerge with two byte arrays of different size needs to return an error and no value")
	}
}

func TestByteSliceExclusiveMergeError(t *testing.T) {
	val, err := ByteSlice([]byte{0x00, 0x00}).ExclusiveMerge(ByteSlice([]byte{0x00}))
	if err == nil || val != nil {
		t.Errorf("ExclusiveMerge with two byte arrays of different size needs to return an error and no value")
	}
}

var testcasesByteSliceNot = []struct {
	name   string
	data   []byte
	result []byte
}{
	{"not empty array", []byte{0xDA, 0x99, 0xBA}, []byte{0x25, 0x66, 0x45}},
	{"empty array", []byte{}, []byte{}},
}

func TestByteSliceNot(t *testing.T) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceNot {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteSlice(tc.data).Not()
			if !reflect.DeepEqual(val, ByteSlice(tc.result)) {
				t.Errorf("Not(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteSliceNot(b *testing.B) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceNot {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteSlice(tc.data).Not()
			}
		})
	}
}

var testcasesByteSliceLeftPad = []struct {
	name   string
	data   []byte
	length int
	filler byte
	result []byte
}{
	{"lpad to two elements on an smaller array", []byte{0xDA}, 2, 0x00, []byte{0x00, 0xDA}},
	{"lpad to two elements with a filler on an smaller array", []byte{0xDA}, 2, 0x00, []byte{0x00, 0xDA}},
	{"lpad to two elements on an empty array", []byte{}, 2, 0x11, []byte{0x11, 0x11}},
	{"lpad to one element on an larger array", []byte{0xDA, 0x99, 0xBA}, 1, 0x11, []byte{0xDA, 0x99, 0xBA}},
	{"lpad to negative element size change nothing", []byte{0xDA}, -1, 0x11, []byte{0xDA}},
}

func TestByteSliceLeftPad(t *testing.T) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceLeftPad {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteSlice(tc.data).leftPad(tc.length, tc.filler)
			if !reflect.DeepEqual(val, ByteSlice(tc.result)) {
				t.Errorf("ByteSlice(%x).leftPad(%v, %v) was %x, should be %x",
					tc.data, tc.length, tc.filler,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteSliceLeftPad(b *testing.B) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceLeftPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteSlice(tc.data).leftPad(tc.length, tc.filler)
			}
		})
	}
}

var testcasesByteSliceRightPad = []struct {
	name   string
	data   []byte
	length int
	filler byte
	result []byte
}{
	{"rpad to two elements on an smaller array", []byte{0xDA}, 2, 0x00, []byte{0xDA, 0x00}},
	{"lpad to two elements with a filler on an smaller array", []byte{0xDA}, 2, 0x00, []byte{0xDA, 0x00}},
	{"rpad to two elements on an empty array", []byte{}, 2, 0x11, []byte{0x11, 0x11}},
	{"rpad to one element on an larger array", []byte{0xDA, 0x99, 0xBA}, 1, 0x11, []byte{0xDA, 0x99, 0xBA}},
	{"rpad to negative element size change nothing", []byte{0xDA}, -1, 0x00, []byte{0xDA}},
}

func TestByteSliceRightPad(t *testing.T) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceRightPad {
		t.Run(tc.name, func(t *testing.T) {
			val = ByteSlice(tc.data).rightPad(tc.length, tc.filler)
			if !reflect.DeepEqual(val, ByteSlice(tc.result)) {
				t.Errorf("ByteSlice(%x).rightPad(%v, %v) was %x, should be %x",
					tc.data, tc.length, tc.filler,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkByteSliceRightPad(b *testing.B) {
	var val ByteSlice
	for _, tc := range testcasesByteSliceRightPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = ByteSlice(tc.data).rightPad(tc.length, tc.filler)
			}
		})
	}
}
