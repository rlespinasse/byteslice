package byteslice

import (
	"errors"
	"math"
)

// Reverse change the order of the byte slice.
func Reverse(data []byte) []byte {
	if len(data) < 2 {
		return data
	}
	sliceLength := len(data)
	sliceHalfLength := int(math.Floor(float64(sliceLength / 2)))
	reversedSlice := make([]byte, sliceLength)

	for i := 0; i <= sliceHalfLength; i++ {
		reversedSlice[i] = data[sliceLength-1-i]
		reversedSlice[sliceLength-1-i] = data[i]
	}
	return reversedSlice
}

// LShift apply left shift operation to an byte slice.
func LShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	dataLength := len(data)
	result := make([]byte, dataLength)
	if shift > maxBitsLength {
		copy(result, data[1:])
		result = LShift(result, shift-maxBitsLength)
	} else {
		for i := dataLength - 1; i >= 0; i-- {
			if i > 0 {
				result[i-1] = data[i] >> (maxBitsLength - shift)
			}
			result[i] = result[i] | (data[i] << shift)
		}
	}
	return result
}

// RShift apply right shift operation to an byte slice.
func RShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	dataLength := len(data)
	result := make([]byte, dataLength)
	if shift > maxBitsLength {
		shiftedData := append(make([]byte, 1), data[:dataLength-1]...)
		result = RShift(shiftedData, shift-maxBitsLength)
	} else {
		for i := 0; i < dataLength; i++ {
			if i < dataLength-1 {
				result[i+1] = data[i] << (maxBitsLength - shift)
			}
			result[i] = result[i] | (data[i] >> shift)
		}
	}
	return result
}

// LPad pads the left-side of a byte slice with a filler byte.
func LPad(data []byte, length int, filler byte) []byte {
	dataLength := len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	result := make([]byte, length-dataLength)
	for i := range result {
		result[i] = filler
	}
	result = append(result, data...)
	return result
}

// RPad pads the right-side of a byte slice with a filler byte.
func RPad(data []byte, length int, filler byte) []byte {
	dataLength := len(data)
	if length < 1 || length <= dataLength {
		return data
	}
	result := make([]byte, length-dataLength)
	for i := range result {
		result[i] = filler
	}
	result = append(data, result...)
	return result
}

// Unset apply AND operation on a byte slice with an "unset" byte slice (must have the same size).
func Unset(data, unsetData []byte) ([]byte, error) {
	var dataLength = len(data)
	var unsetDataLength = len(unsetData)
	if dataLength != unsetDataLength {
		return nil, errors.New("data and unsetData must have the same size")
	}

	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] & unsetData[i]
	}
	return result, nil
}

// Set apply OR operation on a byte slice with an "set" byte slice (must have the same size).
func Set(data, setData []byte) ([]byte, error) {
	dataLength := len(data)
	setDataLength := len(setData)
	if dataLength != setDataLength {
		return nil, errors.New("data and setData must have the same size")
	}

	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] | setData[i]
	}
	return result, nil
}

// Toogle apply XOR operation on a byte slice with an "toogle" byte slice (must have the same size).
func Toogle(data, toogleData []byte) ([]byte, error) {
	dataLength := len(data)
	toogleDataLength := len(toogleData)
	if dataLength != toogleDataLength {
		return nil, errors.New("data and toogleData must have the same size")
	}

	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = data[i] ^ toogleData[i]
	}
	return result, nil
}

// Flip apply NOT operation to a byte slice to flip it.
func Flip(data []byte) []byte {
	dataLength := len(data)
	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		result[i] = ^data[i]
	}
	return result
}

func computeSize(leastSignificantBit, mostSignificantBit uint64) uint64 {
	count := float64(mostSignificantBit-leastSignificantBit) / float64(maxBitsLength)
	return uint64(math.Ceil(count))
}
