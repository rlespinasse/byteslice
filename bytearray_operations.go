package gobits

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

// ExtractBytes get a byte array as subset of another byte array
func ExtractBytes(data []byte, lsbPosition, msbPosition uint64) []byte {
	var maxMsb = uint64(byteLength*len(data) - 1)

	var result = RightShift(data, maxMsb-msbPosition)
	var correctiveShift = maxMsb + lsbPosition - msbPosition
	result = LeftShift(result, correctiveShift)

	var size = ComputeSize(lsbPosition, msbPosition)
	return Trim(result, size)
}

// ComputeSize compute size from 2 bit positions
func ComputeSize(lsbPosition, msbPosition uint64) uint64 {
	var byteCount = float64(msbPosition-lsbPosition) / float64(byteLength)
	return uint64(math.Ceil(byteCount))
}

// Trim get a trim byte array as subset of another byte array
func Trim(data []byte, newSize uint64) []byte {
	return data[uint64(len(data))-newSize:]
}
