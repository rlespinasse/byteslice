package byteslice

import (
	"reflect"
	"testing"
)

var tcLUnset = []struct {
	name      string
	data      []byte
	unsetData []byte
	result    []byte
}{
	{"data and unsetData of equal length", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data shorter than unsetData", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x88, 0x11, 0xAA}, []byte{0x88, 0x11, 0xAA}},
	{"data longer than unsetData", []byte{0xAD, 0x11, 0xAB, 0x88, 0x11, 0xAA}, []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0x88, 0x11, 0xAA}},
	{"empty unsetData on data", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"unsetData on empty data", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{}},
	{"empty unsetData on empty data", []byte{}, []byte{}, []byte{}},
}

func TestLUnset(t *testing.T) {
	var val []byte
	for _, tc := range tcLUnset {
		t.Run(tc.name, func(t *testing.T) {
			val = LUnset(tc.data, tc.unsetData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LUnset(%x, %x) was %x, should be %x",
					tc.data, tc.unsetData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLUnset(b *testing.B) {
	var val []byte
	for _, tc := range tcLUnset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LUnset(tc.data, tc.unsetData)
			}
		})
	}
}

var tcLSet = []struct {
	name    string
	data    []byte
	setData []byte
	result  []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"data longer", []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"setData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"setData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestLSet(t *testing.T) {
	var val []byte
	for _, tc := range tcLSet {
		t.Run(tc.name, func(t *testing.T) {
			val = LSet(tc.data, tc.setData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LSet(%x, %x) was %x, should be %x",
					tc.data, tc.setData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLSet(b *testing.B) {
	var val []byte
	for _, tc := range tcLSet {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LSet(tc.data, tc.setData)
			}
		})
	}
}

var tcLToogle = []struct {
	name       string
	data       []byte
	toogleData []byte
	result     []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"data longer", []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"toogleData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"toogleData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestLToogle(t *testing.T) {
	var val []byte
	for _, tc := range tcLToogle {
		t.Run(tc.name, func(t *testing.T) {
			val = LToogle(tc.data, tc.toogleData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LToogle(%x, %x) was %x, should be %x",
					tc.data, tc.toogleData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLToogle(b *testing.B) {
	var val []byte
	for _, tc := range tcLToogle {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LToogle(tc.data, tc.toogleData)
			}
		})
	}
}

var tcLSubset = []struct {
	name                                    string
	data                                    []byte
	leastSignificantBit, mostSignificantBit uint64
	result                                  []byte
}{
	{"extract nothing", []byte{0xDA, 0x99, 0xBA}, 0, 0, []byte{}},
	{"extract nothing due to inversed positions", []byte{0xDA, 0x99, 0xBA}, 16, 8, []byte{}},
	{"extract nothing due to wrong positions", []byte{0xDA, 0x99, 0xBA}, 100, 101, []byte{}},
	{"extract only in one byte", []byte{0xDA, 0x99, 0xBA}, 5, 7, []byte{0x40}},
	{"extract one byte over two bytes", []byte{0xDA, 0x99, 0xBA}, 7, 8, []byte{0x40}},
	{"extract two bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 6, 17, []byte{0xA6, 0x60}},
	{"extract three bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 1, 22, []byte{0xB5, 0x33, 0x74}},
	{"extract all bytes", []byte{0xDA, 0x99, 0xBA}, 0, 23, []byte{0xDA, 0x99, 0xBA}},
	{"extract all bytes with an overflow position", []byte{0xDA, 0x99, 0xBA}, 0, 100, []byte{0xDA, 0x99, 0xBA}},
}

func TestLSubset(t *testing.T) {
	var val []byte
	for _, tc := range tcLSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = LSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LSubset(%x, %v, %v) was %x, should be %x",
					tc.data, tc.leastSignificantBit, tc.mostSignificantBit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLSubset(b *testing.B) {
	var val []byte
	for _, tc := range tcLSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			}
		})
	}
}
