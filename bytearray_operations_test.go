package gobits

import (
	"reflect"
	"testing"
)

func TestRightShift(t *testing.T) {
	var result = RightShift([]byte{0x99, 0xBA}, 1)
	if !reflect.DeepEqual(result, []byte{0x33, 0x74}) {
		t.Errorf("RightShift don't work: %x", result)
	}
}

func TestLeftShift(t *testing.T) {
	var result = LeftShift([]byte{0x99, 0xBA}, 1)
	if !reflect.DeepEqual(result, []byte{0x4C, 0xDD}) {
		t.Errorf("LeftShift don't work: %x", result)
	}
}
