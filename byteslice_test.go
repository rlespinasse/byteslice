package byteslice

import (
	"reflect"
	"testing"
)

var tcReverse = []struct {
	name   string
	data   []byte
	result []byte
}{
	{"empty byte slice", []byte{}, []byte{}},
	{"one element byte slice", []byte{0x01}, []byte{0x01}},
	{"even length byte slice", []byte{0x01, 0x10}, []byte{0x10, 0x01}},
	{"odd length byte slice", []byte{0x01, 0xFF, 0x10}, []byte{0x10, 0xFF, 0x01}},
}

func TestReverse(t *testing.T) {
	var val []byte
	for _, tc := range tcReverse {
		t.Run(tc.name, func(t *testing.T) {
			val = Reverse(tc.data)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("Reverse(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	var val []byte
	for _, tc := range tcReverse {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = Reverse(tc.data)
			}
		})
	}
}

var tcLShift = []struct {
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

func TestLShift(t *testing.T) {
	var val []byte
	for _, tc := range tcLShift {
		t.Run(tc.name, func(t *testing.T) {
			val = LShift(tc.data, tc.shift)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLShift(b *testing.B) {
	var val []byte
	for _, tc := range tcLShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LShift(tc.data, tc.shift)
			}
		})
	}
}

var tcRShift = []struct {
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

func TestRShift(t *testing.T) {
	var val []byte
	for _, tc := range tcRShift {
		t.Run(tc.name, func(t *testing.T) {
			val = RShift(tc.data, tc.shift)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RShift(%x, %v) was %x, should be %x",
					tc.data, tc.shift,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRShift(b *testing.B) {
	var val []byte
	for _, tc := range tcRShift {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RShift(tc.data, tc.shift)
			}
		})
	}
}

var tcLPad = []struct {
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

func TestLPad(t *testing.T) {
	var val []byte
	for _, tc := range tcLPad {
		t.Run(tc.name, func(t *testing.T) {
			val = LPad(tc.data, tc.length, tc.filler)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LPad(%x, %v, %v) was %x, should be %x",
					tc.data, tc.length, tc.filler,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLPad(b *testing.B) {
	var val []byte
	for _, tc := range tcLPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LPad(tc.data, tc.length, tc.filler)
			}
		})
	}
}

var tcRPad = []struct {
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

func TestRPad(t *testing.T) {
	var val []byte
	for _, tc := range tcRPad {
		t.Run(tc.name, func(t *testing.T) {
			val = RPad(tc.data, tc.length, tc.filler)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RPad(%x, %v, %v) was %x, should be %x",
					tc.data, tc.length, tc.filler,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRPad(b *testing.B) {
	var val []byte
	for _, tc := range tcRPad {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RPad(tc.data, tc.length, tc.filler)
			}
		})
	}
}

func TestUnsetError(t *testing.T) {
	val, err := Unset([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("Unset with two byte slices of different size needs to return an error and no value")
	}
}

func TestSetError(t *testing.T) {
	val, err := Set([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("Set with two byte slices of different size needs to return an error and no value")
	}
}

func TestToogleError(t *testing.T) {
	val, err := Toogle([]byte{0x00, 0x00}, []byte{0x00})
	if err == nil || val != nil {
		t.Errorf("Toogle with two byte slices of different size needs to return an error and no value")
	}
}

var tcFlip = []struct {
	name   string
	data   []byte
	result []byte
}{
	{"not empty array", []byte{0xDA, 0x99, 0xBA}, []byte{0x25, 0x66, 0x45}},
	{"empty array", []byte{}, []byte{}},
}

func TestFlip(t *testing.T) {
	var val []byte
	for _, tc := range tcFlip {
		t.Run(tc.name, func(t *testing.T) {
			val = Flip(tc.data)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("Flip(%x) was %x, should be %x",
					tc.data,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkFlip(b *testing.B) {
	var val []byte
	for _, tc := range tcFlip {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = Flip(tc.data)
			}
		})
	}
}
