package byteness

import (
	"reflect"
	"testing"
)

var testcasesLittleEndianOperationsMask = []struct {
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

func TestLittleEndianOperationsMask(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsMask {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianOperations.Mask(tc.data, tc.mask)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianOperations.Mask(%x, %x) was %x, should be %x",
					tc.data, tc.mask,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianOperationsMask(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsMask {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianOperations.Mask(tc.data, tc.mask)
			}
		})
	}
}

var testcasesLittleEndianOperationsInclusiveMerge = []struct {
	name   string
	data1  []byte
	data2  []byte
	result []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB}},
	{"first array longer", []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0xFF, 0x99, 0xBB, 0xAD, 0x11, 0xAB}, []byte{0xFF, 0x99, 0xBB, 0xFF, 0x99, 0xBB}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestLittleEndianOperationsInclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsInclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianOperations.InclusiveMerge(tc.data1, tc.data2)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianOperations.InclusiveMerge(%x, %x) was %x, should be %x",
					tc.data1, tc.data2,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianOperationsInclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsInclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianOperations.InclusiveMerge(tc.data1, tc.data2)
			}
		})
	}
}

var testcasesLittleEndianOperationsExclusiveMerge = []struct {
	name   string
	data1  []byte
	data2  []byte
	result []byte
}{
	{"equal length arrays", []byte{0xDA, 0x99, 0xBA}, []byte{0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11}},
	{"first array longer", []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"second array longer", []byte{0xDA, 0x99, 0xBA}, []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}, []byte{0x77, 0x88, 0x11, 0x77, 0x88, 0x11}},
	{"first array empty", []byte{}, []byte{0xAD, 0x11, 0xAB}, []byte{0xAD, 0x11, 0xAB}},
	{"second array empty", []byte{0xDA, 0x99, 0xBA}, []byte{}, []byte{0xDA, 0x99, 0xBA}},
	{"empty arrays", []byte{}, []byte{}, []byte{}},
}

func TestLittleEndianOperationsExclusiveMerge(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsExclusiveMerge {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianOperations.ExclusiveMerge(tc.data1, tc.data2)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianOperations.ExclusiveMerge(%x, %x) was %x, should be %x",
					tc.data1, tc.data2,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianOperationsExclusiveMerge(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsExclusiveMerge {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianOperations.ExclusiveMerge(tc.data1, tc.data2)
			}
		})
	}
}

var testcasesLittleEndianOperationsExtractBytes = []struct {
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

func TestLittleEndianOperationsExtractBytes(t *testing.T) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsExtractBytes {
		t.Run(tc.name, func(t *testing.T) {
			val = LittleEndianOperations.ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			if !reflect.DeepEqual(val, tc.result) {
				t.Errorf("LittleEndianOperations.ExtractBytes(%x, %v, %v) was %x, should be %x",
					tc.data, tc.lsbPosition, tc.msbPosition,
					val,
					tc.result)
			}
		})
	}
}

func BenchmarkLittleEndianOperationsExtractBytes(b *testing.B) {
	var val []byte
	for _, tc := range testcasesLittleEndianOperationsExtractBytes {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val = LittleEndianOperations.ExtractBytes(tc.data, tc.lsbPosition, tc.msbPosition)
			}
		})
	}
}
