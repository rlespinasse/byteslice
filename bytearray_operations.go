package byteness

import "math"

// RightShift apply right shift operation to byte array data
func RightShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make([]byte, dataLength)
	if shift > byteLength {
		copy(result, data[1:])
		result = RightShift(result, shift-byteLength)
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

// LeftShift apply left shift operation to byte array data
func LeftShift(data []byte, shift uint64) []byte {
	if shift == 0 {
		return data
	}
	var dataLength = len(data)
	result := make([]byte, dataLength)
	if shift > byteLength {
		var shiftedData = append(make([]byte, 1), data[:dataLength-1]...)
		result = LeftShift(shiftedData, shift-byteLength)
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

// Mask apply AND mask to byte array
func Mask(data, mask []byte) []byte {
	var dataLength = len(data)
	if dataLength < 1 {
		return data
	}
	result := make([]byte, dataLength)
	copy(result, data)

	var maskLength = len(mask)
	var operationLength = maskLength
	// If mask is longer than data, keep operation to (shorter) data length
	if maskLength > dataLength {
		operationLength = dataLength
	}
	for i := 1; i < operationLength+1; i++ {
		result[dataLength-i] &= mask[maskLength-i]
	}
	return result
}

// InclusiveMerge two byte arrays
func InclusiveMerge(data1, data2 []byte) []byte {
	var shorterData = data1
	var longerData = data2
	var longerLength = len(data2)
	if longerLength < len(data1) {
		shorterData = data2
		longerData = data1
		longerLength = len(data1)
	}
	result := make([]byte, longerLength)
	copy(result[longerLength-len(shorterData):], shorterData)

	for i := longerLength - 1; i > -1; i-- {
		result[i] |= longerData[i]
	}
	return result
}

// ExclusiveMerge two byte arrays
func ExclusiveMerge(data1, data2 []byte) []byte {
	var shorterData = data1
	var longerData = data2
	var longerLength = len(data2)
	if longerLength < len(data1) {
		shorterData = data2
		longerData = data1
		longerLength = len(data1)
	}
	result := make([]byte, longerLength)
	copy(result[longerLength-len(shorterData):], shorterData)

	for i := longerLength - 1; i > -1; i-- {
		result[i] ^= longerData[i]
	}
	return result
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

// ExtractBytes get a byte array as subset of another byte array
func ExtractBytes(data []byte, lsbPosition, msbPosition uint64) []byte {
	var maxMsb = uint64(byteLength*len(data) - 1)

	if msbPosition <= lsbPosition || lsbPosition > maxMsb {
		return make([]byte, 0)
	}

	if msbPosition > maxMsb {
		msbPosition = maxMsb
	}

	var result = RightShift(data, maxMsb-msbPosition)
	var correctiveShift = maxMsb + lsbPosition - msbPosition
	result = LeftShift(result, correctiveShift)

	var size = computeSize(lsbPosition, msbPosition)
	return trim(result, size)
}

func computeSize(lsbPosition, msbPosition uint64) uint64 {
	var byteCount = float64(msbPosition-lsbPosition) / float64(byteLength)
	return uint64(math.Ceil(byteCount))
}

func trim(data []byte, newSize uint64) []byte {
	return data[uint64(len(data))-newSize:]
}
