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
	for _, tc := range tcLUnset {
		t.Run(tc.name, func(t *testing.T) {
			values = LUnset(tc.data, tc.unsetData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("LUnset(%x, %x) was %x, should be %x",
					tc.data, tc.unsetData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkLUnset(b *testing.B) {
	for _, tc := range tcLUnset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = LUnset(tc.data, tc.unsetData)
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
	for _, tc := range tcLSet {
		t.Run(tc.name, func(t *testing.T) {
			values = LSet(tc.data, tc.setData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("LSet(%x, %x) was %x, should be %x",
					tc.data, tc.setData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkLSet(b *testing.B) {
	for _, tc := range tcLSet {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = LSet(tc.data, tc.setData)
			}
		})
	}
}

var tcLToggle = []struct {
	name       string
	data       []byte
	toggleData []byte
	result     []byte
}{
	{"equal length slices", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"data longer", []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"toggleData longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"data empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"toggleData empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty slices", []byte{}, []byte{}, []byte{}},
}

func TestLToggle(t *testing.T) {
	for _, tc := range tcLToggle {
		t.Run(tc.name, func(t *testing.T) {
			values = LToggle(tc.data, tc.toggleData)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("LToggle(%x, %x) was %x, should be %x",
					tc.data, tc.toggleData,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkLToggle(b *testing.B) {
	for _, tc := range tcLToggle {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = LToggle(tc.data, tc.toggleData)
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
	for _, tc := range tcLSubset {
		t.Run(tc.name, func(t *testing.T) {
			values = LSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			if !reflect.DeepEqual(values, tc.result) {
				t.Errorf("LSubset(%x, %v, %v) was %x, should be %x",
					tc.data, tc.leastSignificantBit, tc.mostSignificantBit,
					values,
					tc.result)
			}
		})
	}
}

func BenchmarkLSubset(b *testing.B) {
	for _, tc := range tcLSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				values = LSubset(tc.data, tc.leastSignificantBit, tc.mostSignificantBit)
			}
		})
	}
}
