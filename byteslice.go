package byteslice

import (
	"errors"
	"math"
)

// ByteSlice attaches the methods of Interface to []byte without specific ordering
type ByteSlice []byte

// Reverse reverse the byte slice
func (data ByteSlice) Reverse() ByteSlice {
	var sliceLength = len(data)
	var reversedSlice = make(ByteSlice, sliceLength)

	for i := range data {
		reversedSlice[sliceLength-i-1] = data[i]
	}
	return reversedSlice
}

// LeftShift apply left shift operation to byte array data
func (data ByteSlice) LeftShift(shift uint64) ByteSlice {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make(ByteSlice, dataLength)
	if shift > byteLength {
		copy(result, data[1:])
		result = ByteSlice(result).LeftShift(shift - byteLength)
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
func (data ByteSlice) RightShift(shift uint64) ByteSlice {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make(ByteSlice, dataLength)
	if shift > byteLength {
		var shiftedData = append(make(ByteSlice, 1), data[:dataLength-1]...)
		result = ByteSlice(shiftedData).RightShift(shift - byteLength)
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
func (data ByteSlice) Mask(mask ByteSlice) (ByteSlice, error) {
	var dataLength = len(data)
	var maskLength = len(mask)
	if dataLength != maskLength {
		return nil, errors.New("data and mask must have the same size")
	}

	result := make(ByteSlice, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] & mask[i]
	}
	return result, nil
}

// InclusiveMerge apply OR operation on two byte arrays (must have the same size)
func (data ByteSlice) InclusiveMerge(anotherData ByteSlice) (ByteSlice, error) {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)
	if dataLength != anotherDataLength {
		return nil, errors.New("data and anotherData must have the same size")
	}

	result := make(ByteSlice, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] | anotherData[i]
	}
	return result, nil
}

// ExclusiveMerge apply XOR operation on two byte arrays (must have the same size)
func (data ByteSlice) ExclusiveMerge(anotherData ByteSlice) (ByteSlice, error) {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)
	if dataLength != anotherDataLength {
		return nil, errors.New("data and anotherData must have the same size")
	}

	result := make(ByteSlice, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] ^ anotherData[i]
	}
	return result, nil
}

// Not apply NOT operation to byte array
func (data ByteSlice) Not() ByteSlice {
	var dataLength = len(data)
	result := make(ByteSlice, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = ^data[i]
	}
	return result
}

func (data ByteSlice) leftPad(length int, filler byte) ByteSlice {
	var dataLength = len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	var result = make(ByteSlice, length-dataLength, length)
	for i := range result {
		result[i] = filler
	}
	result = append(result, data...)
	return result
}

func (data ByteSlice) rightPad(length int, filler byte) ByteSlice {
	var dataLength = len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	var result = make(ByteSlice, length-dataLength, length)
	for i := range result {
		result[i] = filler
	}
	result = append(data, result...)
	return result
}

func computeSize(lsbPosition, msbPosition uint64) uint64 {
	var byteCount = float64(msbPosition-lsbPosition) / float64(byteLength)
	return uint64(math.Ceil(byteCount))
}
