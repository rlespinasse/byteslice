package byteness

import (
	"errors"
	"math"
)

// LeftShift apply left shift operation to byte array data
func LeftShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make([]byte, dataLength)
	if shift > byteLength {
		copy(result, data[1:])
		result = LeftShift(result, shift-byteLength)
	} else {
		for i := dataLength - 1; i >= 0; i-- {
			if i > 0 {
				result[i-1] = data[i] >> (byteLength - shift)
			}
			result[i] = result[i] | (data[i] << shift)
		}
	}
	return result
}

// RightShift apply right shift operation to byte array data
func RightShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make([]byte, dataLength)
	if shift > byteLength {
		var shiftedData = append(make([]byte, 1), data[:dataLength-1]...)
		result = RightShift(shiftedData, shift-byteLength)
	} else {
		for i := 0; i < dataLength; i++ {
			if i < dataLength-1 {
				result[i+1] = data[i] << (byteLength - shift)
			}
			result[i] = result[i] | (data[i] >> shift)
		}
	}
	return result
}

// Mask apply AND mask to a byte array, data and mask must have the same size
func Mask(data, mask []byte) ([]byte, error) {
	var dataLength = len(data)
	var maskLength = len(mask)
	if dataLength != maskLength {
		return nil, errors.New("data and mask must have the same size")
	}

	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] & mask[i]
	}
	return result, nil
}

// InclusiveMerge apply OR operation on two byte arrays (must have the same size)
func InclusiveMerge(data1, data2 []byte) ([]byte, error) {
	var data1Length = len(data1)
	var data2Length = len(data2)
	if data1Length != data2Length {
		return nil, errors.New("data1 and data2 must have the same size")
	}

	result := make([]byte, data1Length)
	for i := 0; i < data1Length; i++ {
		result[i] = data1[i] | data2[i]
	}
	return result, nil
}

// ExclusiveMerge apply XOR operation on two byte arrays (must have the same size)
func ExclusiveMerge(data1, data2 []byte) ([]byte, error) {
	var data1Length = len(data1)
	var data2Length = len(data2)
	if data1Length != data2Length {
		return nil, errors.New("data1 and data2 must have the same size")
	}

	result := make([]byte, data1Length)
	for i := 0; i < data1Length; i++ {
		result[i] = data1[i] ^ data2[i]
	}
	return result, nil
}

// Not apply NOT operation to byte array
func Not(data []byte) []byte {
	var dataLength = len(data)
	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = ^data[i]
	}
	return result
}

func computeSize(lsbPosition, msbPosition uint64) uint64 {
	var byteCount = float64(msbPosition-lsbPosition) / float64(byteLength)
	return uint64(math.Ceil(byteCount))
}

func leftPad(data []byte, length int, filler byte) []byte {
	var dataLength = len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	var result = make([]byte, length-dataLength, length)
	for i := range result {
		result[i] = filler
	}
	result = append(result, data...)
	return result
}

func rightPad(data []byte, length int, filler byte) []byte {
	var dataLength = len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	var result = make([]byte, length-dataLength, length)
	for i := range result {
		result[i] = filler
	}
	result = append(data, result...)
	return result
}
