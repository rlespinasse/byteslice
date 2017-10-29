package byteslice

import (
	"reflect"
	"testing"
)

var tcRUnset = []struct {
	name      string
	data      []byte
	unsetData []byte
	result    []byte
}{
	{"data and unsetData of equal length", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data shorter than unsetData", []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data longer than unsetData", []byte{0x88, 0x11, 0xAA, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0x88, 0x11, 0xAA}},
	{"empty unsetData on data", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"unsetData on empty data", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{}},
	{"empty unsetData on empty data", []byte{}, []byte{}, []byte{}},
}

func TestRUnset(t *testing.T) {
	var val []byte
	for _, tc := range tcRUnset {
		t.Run(tc.name, func(t *testing.T) {
			val = RUnset(tc.data, tc.unsetData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RUnset(%x, %x) was %x, should be %x",
					tc.data, tc.unsetData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRUnset(b *testing.B) {
	var val []byte
	for _, tc := range tcRUnset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RUnset(tc.data, tc.unsetData)
			}
		})
	}
}

var tcRSet = []struct {
	name    string
	data    []byte
	setData []byte
	result  []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"data longer", []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"setData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"setData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestRSet(t *testing.T) {
	var val []byte
	for _, tc := range tcRSet {
		t.Run(tc.name, func(t *testing.T) {
			val = RSet(tc.data, tc.setData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RSet(%x, %x) was %x, should be %x",
					tc.data, tc.setData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRSet(b *testing.B) {
	var val []byte
	for _, tc := range tcRSet {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RSet(tc.data, tc.setData)
			}
		})
	}
}

var tcRToogle = []struct {
	name       string
	data       []byte
	toogleData []byte
	result     []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"data longer", []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"toogleData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"toogleData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestRToogle(t *testing.T) {
	var val []byte
	for _, tc := range tcRToogle {
		t.Run(tc.name, func(t *testing.T) {
			val = RToogle(tc.data, tc.toogleData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RToogle(%x, %x) was %x, should be %x",
					tc.data, tc.toogleData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRToogle(b *testing.B) {
	var val []byte
	for _, tc := range tcRToogle {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RToogle(tc.data, tc.toogleData)
			}
		})
	}
}

var tcRSubset = []struct {
	name                                    string
	data                                    []byte
	leastSignificantBit, mostSignificantBit uint64
	result                                  []byte
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

func TestRSubset(t *testing.T) {
	var val []byte
	for _, tc := range tcRSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = RSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("RSubset(%x, %v, %v) was %x, should be %x",
					tc.data, tc.leastSignificantBit, tc.mostSignificantBit,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkRSubset(b *testing.B) {
	var val []byte
	for _, tc := range tcRSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = RSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			}
		})
	}
}
