package byteness

import (
	"reflect"
	"testing"
)

var testcasesBigEndianOperationsMask = []struct {
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

func TestBigEndianOperationsMask(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsMask {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianOperations.Mask(tc.data, tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianOperations.Mask(%x, %x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianOperationsMask(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsMask {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianOperations.Mask(tc.data, tc.mask)
			}
		})
	}
}

var testcasesBigEndianOperationsInclusiveMerge = []struct {
	name   string
	data1  []byte
	data2  []byte
	result []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"first array longer", []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0xFF, 0x99, 0xBB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestBigEndianOperationsInclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsInclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianOperations.InclusiveMerge(tc.data1, tc.data2)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianOperations.InclusiveMerge(%x, %x) was %x, should be %x",
					tc.data1, tc.data2,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianOperationsInclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsInclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianOperations.InclusiveMerge(tc.data1, tc.data2)
			}
		})
	}
}

var testcasesBigEndianOperationsExclusiveMerge = []struct {
	name   string
	data1  []byte
	data2  []byte
	result []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"first array longer", []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB, 0x77, 0x88, 0x11}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestBigEndianOperationsExclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsExclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianOperations.ExclusiveMerge(tc.data1, tc.data2)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianOperations.ExclusiveMerge(%x, %x) was %x, should be %x",
					tc.data1, tc.data2,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianOperationsExclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsExclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianOperations.ExclusiveMerge(tc.data1, tc.data2)
			}
		})
	}
}

var testcasesBigEndianOperationsExtractBytes = []struct {
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

func TestBigEndianOperationsExtractBytes(t *testing.T) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsExtractBytes {
		t.Run(tc.name, func(t *testing.T) {
			val = BigEndianOperations.ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("BigEndianOperations.ExtractBytes(%x, %v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkBigEndianOperationsExtractBytes(b *testing.B) {
	var val []byte
	for _, tc := range testcasesBigEndianOperationsExtractBytes {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = BigEndianOperations.ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
