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
	for _, tc := range tcRUnset {
		t.Run(tc.name, func(t *testing.T) {
			values = RUnset(tc.data, tc.unsetData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("RUnset(%x, %x) was %x, should be %x",
					tc.data, tc.unsetData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkRUnset(b *testing.B) {
	for _, tc := range tcRUnset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = RUnset(tc.data, tc.unsetData)
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
	for _, tc := range tcRSet {
		t.Run(tc.name, func(t *testing.T) {
			values = RSet(tc.data, tc.setData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("RSet(%x, %x) was %x, should be %x",
					tc.data, tc.setData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkRSet(b *testing.B) {
	for _, tc := range tcRSet {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = RSet(tc.data, tc.setData)
			}
		})
	}
}

var tcRToggle = []struct {
	name       string
	data       []byte
	toggleData []byte
	result     []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"data longer", []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"toggleData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"toggleData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestRToggle(t *testing.T) {
	for _, tc := range tcRToggle {
		t.Run(tc.name, func(t *testing.T) {
			values = RToggle(tc.data, tc.toggleData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("RToggle(%x, %x) was %x, should be %x",
					tc.data, tc.toggleData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkRToggle(b *testing.B) {
	for _, tc := range tcRToggle {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = RToggle(tc.data, tc.toggleData)
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
	for _, tc := range tcRSubset {
		t.Run(tc.name, func(t *testing.T) {
			values = RSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("RSubset(%x, %v, %v) was %x, should be %x",
					tc.data, tc.leastSignificantBit, tc.mostSignificantBit,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkRSubset(b *testing.B) {
	for _, tc := range tcRSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = RSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			}
		})
	}
}
