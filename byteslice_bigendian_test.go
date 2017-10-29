package byteslice

import (
	"reflect"
	"testing"
)

var testcasesBigEndianByteSliceMask = []struct {
	name   string
	data   []byte
	mask   []byte
	result []byte
}{
	{"data and mask of equal length", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x88, 0x11, 0xAA}},
	{"data shorter than mask", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x88, 0x11, 0xAA}, []byte{0x88, 0x11, 0xAA}},
	{"data longer than mask", []byte{0xAD, 0x11, 0xAB, 0x88, 0x11, 0xAA}, []byte{0xDA, 0x99, 0xBA}, []byte{0x88, 0x11, 0xAA, 0x88, 0x11, 0xAA}},
	{"empty mask on data", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"mask on empty data", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{}},
	{"empty mask on empty data", []byte{}, []byte{}, []byte{}},
}

func TestBigEndianByteSliceMask(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceMask {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianByteSlice(tc.data).Mask(tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianByteSlice(%x).Mask(%x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianByteSliceMask(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceMask {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianByteSlice(tc.data).Mask(tc.mask)
			}
		})
	}
}

var testcasesBigEndianByteSliceInclusiveMerge = []struct {
	name        string
	data        []byte
	anotherData []byte
	result      []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"first array longer", []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestBigEndianByteSliceInclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceInclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianByteSlice(tc.data).InclusiveMerge(tc.anotherData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianByteSlice(%x).InclusiveMerge(%x) was %x, should be %x",
					tc.data, tc.anotherData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianByteSliceInclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceInclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianByteSlice(tc.data).InclusiveMerge(tc.anotherData)
			}
		})
	}
}

var testcasesBigEndianByteSliceExclusiveMerge = []struct {
	name        string
	data        []byte
	anotherData []byte
	result      []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"first array longer", []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestBigEndianByteSliceExclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceExclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianByteSlice(tc.data).ExclusiveMerge(tc.anotherData)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianByteSlice(%x).ExclusiveMerge(%x) was %x, should be %x",
					tc.data, tc.anotherData,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianByteSliceExclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceExclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianByteSlice(tc.data).ExclusiveMerge(tc.anotherData)
			}
		})
	}
}

var testcasesBigEndianByteSliceSubset = []struct {
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
	{"extract two bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 6, 17, []byte{0xA6, 0x60}},
	{"extract three bytes over three bytes", []byte{0xDA, 0x99, 0xBA}, 1, 22, []byte{0xB5, 0x33, 0x74}},
	{"extract all bytes", []byte{0xDA, 0x99, 0xBA}, 0, 23, []byte{0xDA, 0x99, 0xBA}},
	{"extract all bytes with an overflow position", []byte{0xDA, 0x99, 0xBA}, 0, 100, []byte{0xDA, 0x99, 0xBA}},
}

func TestBigEndianByteSliceSubset(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianByteSlice(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianByteSlice(%x).Subset(%v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianByteSliceSubset(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianByteSlice(tc.data).Subset(tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}

var testcasesBigEndianByteSliceLittleEndianSubset = []struct {
	name                     string
	data                     []byte
	lsbPosition, msbPosition uint64
	result                   []byte
}{
	{"extract nothing", []byte{0xDA, 0x99, 0xBA}, 0, 0, []byte{}},
	{"extract nothing due to inversed positions", []byte{0xDA, 0x99, 0xBA}, 16, 8, []byte{}},
	{"extract nothing due to wrong positions", []byte{0xDA, 0x99, 0xBA}, 100, 101, []byte{}},
	{"extract only in one byte", []byte{0xBA, 0x99, 0xDA}, 5, 7, []byte{0x05}},
	{"extract one byte over two bytes", []byte{0xBA, 0x99, 0xDA}, 7, 8, []byte{0x03}},
	{"extract two bytes over three bytes", []byte{0xBA, 0x99, 0xDA}, 6, 17, []byte{0x66, 0x0A}},
	{"extract three bytes over three bytes", []byte{0xBA, 0x99, 0xDA}, 1, 22, []byte{0xDD, 0x4C, 0x2D}},
	{"extract all bytes", []byte{0xDA, 0x99, 0xBA}, 0, 23, []byte{0xDA, 0x99, 0xBA}},
	{"extract all bytes with an overflow position", []byte{0xDA, 0x99, 0xBA}, 0, 100, []byte{0xDA, 0x99, 0xBA}},
}

func TestBigEndianByteSliceLittleEndianSubset(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceLittleEndianSubset {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianByteSlice(tc.data).LittleEndianSubset(tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianByteSlice(%x).LittleEndianSubset(%v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianByteSliceLittleEndianSubset(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianByteSliceLittleEndianSubset {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianByteSlice(tc.data).LittleEndianSubset(tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
