package gobits

// RightShift apply right shift operation to byte array data
func RightShift(data []byte, shift uint8) []byte {
	var dataLength = len(data)
	result := make([]byte, dataLength)
	for i := dataLength - 1; i >= 0; i-- {
		if i > 0 {
			result[i-1] = data[i] >> (byteLength - shift)
		}
		result[i] = result[i] | (data[i] << shift)
	}
	return result
}

// LeftShift apply left shift operation to byte array data
func LeftShift(data []byte, shift uint8) []byte {
	var dataLength = len(data)
	result := make([]byte, dataLength)
	for i := 0; i < dataLength; i++ {
		if i < dataLength-1 {
			result[i+1] = data[i] << (byteLength - shift)
		}
		result[i] = result[i] | (data[i] >> shift)
	}
	return result
}
