package byteslice

import (
	"reflect"
	"testing"
)

var testcasesLittleEndianByteSliceMask = []struct {
	name   string
	data   []byte
	mask   []byte
	result []byte
}{
	{"data and mask of equal length", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data shorter than mask", []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data longer than mask", []byte{0x88, 0x11, 0xAA, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0x88, 0x11, 0xAA}},
	{"empty mask on data", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"mask on empty data", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{}},
	{"empty mask on empty data", []byte{}, []byte{}, []byte{}},
}

func TestLittleEndianByteSliceMask(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceMask {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianByteSlice(tc.data).Mask(tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianByteSlice(%x).Mask(%x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianByteSliceMask(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceMask {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianByteSlice(tc.data).Mask(tc.mask)
			}
		})
	}
}

var testcasesLittleEndianByteSliceInclusiveMerge = []struct {
	name        string
	data        []byte
	anotherData []byte
	result      []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"first array longer", []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestLittleEndianByteSliceInclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceInclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianByteSlice(tc.data).InclusiveMerge(tc.anotherData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianByteSlice(%x).InclusiveMerge(%x) was %x, should be %x",
					tc.data, tc.anotherData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianByteSliceInclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceInclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianByteSlice(tc.data).InclusiveMerge(tc.anotherData)
			}
		})
	}
}

var testcasesLittleEndianByteSliceExclusiveMerge = []struct {
	name        string
	data        []byte
	anotherData []byte
	result      []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"first array longer", []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestLittleEndianByteSliceExclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceExclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianByteSlice(tc.data).ExclusiveMerge(tc.anotherData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianByteSlice(%x).ExclusiveMerge(%x) was %x, should be %x",
					tc.data, tc.anotherData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianByteSliceExclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceExclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianByteSlice(tc.data).ExclusiveMerge(tc.anotherData)
			}
		})
	}
}

var testcasesLittleEndianByteSliceSubset = []struct {
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

func TestLittleEndianByteSliceSubset(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianByteSlice(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianByteSlice(%x).Subset(%v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianByteSliceSubset(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianByteSlice(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}

var testcasesLittleEndianByteSliceBigEndianSubset = []struct {
	name                     string
	data                     []byte
	lsbPosition, msbPosition uint64
	result                   []byte
}{
	{"extract nothing", []byte{0xDA, 0x99, 0xBA}, 0, 0, []byte{}},
	{"extract nothing due to inversed positions", []byte{0xDA, 0x99, 0xBA}, 16, 8, []byte{}},
	{"extract nothing due to wrong positions", []byte{0xDA, 0x99, 0xBA}, 100, 101, []byte{}},
	{"extract only in one byte", []byte{0xDA, 0x99, 0xBA}, 5, 7, []byte{0x40}},
	{"extract one byte over two bytes", []byte{0xDA, 0x99, 0xBA}, 7, 8, []byte{0x40}},
	{"extract two bytes over three bytes", []byte{0xBA, 0x99, 0xDA}, 6, 17, []byte{0x60, 0xA6}},
	{"extract three bytes over three bytes", []byte{0xBA, 0x99, 0xDA}, 1, 22, []byte{0x74, 0x33, 0xB5}},
	{"extract all bytes", []byte{0xDA, 0x99, 0xBA}, 0, 23, []byte{0xDA, 0x99, 0xBA}},
	{"extract all bytes with an overflow position", []byte{0xDA, 0x99, 0xBA}, 0, 100, []byte{0xDA, 0x99, 0xBA}},
}

func TestLittleEndianByteSliceBigEndianSubset(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceBigEndianSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianByteSlice(tc.data).BigEndianSubset(tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianByteSlice(%x).BigEndianSubset(%v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianByteSliceBigEndianSubset(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianByteSliceBigEndianSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianByteSlice(tc.data).BigEndianSubset(tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
