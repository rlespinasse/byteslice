package byteness

import (
	"reflect"
	"testing"
)

var testcasesLeftShift = []struct {
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

var testcasesRightShift = []struct {
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

func TestMaskError(t *testing.T) {
	val, err := Mask([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("Mask with two byte arrays of different size needs to return an error and no value")
	}
}

func TestInclusiveMergeError(t *testing.T) {
	val, err := InclusiveMerge([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("InclusiveMerge with two byte arrays of different size needs to return an error and no value")
	}
}

func TestExclusiveMergeError(t *testing.T) {
	val, err := ExclusiveMerge([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("ExclusiveMerge with two byte arrays of different size needs to return an error and no value")
	}
}

var testcasesNot = []struct {
	name   string
	data   []byte
	result []byte
}{
	{"not empty array", []byte{0xDA, 0x99, 0xBA}, []byte{0x25, 0x66, 0x45}},
	{"empty array", []byte{}, []byte{}},
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

func BenchmarkNot(b *testing.B) {
	var val []byte
	for _, tc := range testcasesNot {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = Not(tc.data)
			}
		})
	}
}

var testcasesLeftPad = []struct {
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

func TestLeftPad(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLeftPad {
		t.Run(tc.name, func(t *testing.T) {
			val = leftPad(tc.data, tc.length, tc.filler)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("leftPad(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLeftPad(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLeftPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = leftPad(tc.data, tc.length, tc.filler)
			}
		})
	}
}

var testcasesRightPad = []struct {
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

func TestRightPad(t *testing.T) {
	var val []byte
	for _, tc := range testcasesRightPad {
		t.Run(tc.name, func(t *testing.T) {
			val = rightPad(tc.data, tc.length, tc.filler)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("rightPad(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRightPad(b *testing.B) {
	var val []byte
	for _, tc := range testcasesRightPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = rightPad(tc.data, tc.length, tc.filler)
			}
		})
	}
}
