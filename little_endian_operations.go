package byteness

type littleEndianOperations struct{}

// LittleEndianOperations test
var LittleEndianOperations = &littleEndianOperations{}

// Mask apply AND mask to a byte array
func (op *littleEndianOperations) Mask(data, mask []byte) []byte {
	var dataLength = len(data)
	if dataLength < 1 {
		return data
	}

	var maskLength = len(mask)
	var operationLength = dataLength
	var operationCut = 0
	if maskLength > dataLength {
		operationLength = maskLength
		operationCut = operationLength - dataLength
	}

	result, _ := Mask(leftPad(data, operationLength, 0xFF), leftPad(mask, operationLength, 0xFF))
	return result[operationCut:]
}

// InclusiveMerge apply OR operation on two byte arrays
func (op *littleEndianOperations) InclusiveMerge(data1, data2 []byte) []byte {
	var data1Length = len(data1)
	var data2Length = len(data2)

	var operationLength = data1Length
	if data2Length > data1Length {
		operationLength = data2Length
	}

	result, _ := InclusiveMerge(leftPad(data1, operationLength, 0x00), leftPad(data2, operationLength, 0x00))
	return result
}

// ExclusiveMerge apply XOR operation on two byte arrays
func (op *littleEndianOperations) ExclusiveMerge(data1, data2 []byte) []byte {
	var data1Length = len(data1)
	var data2Length = len(data2)

	var operationLength = data1Length
	if data2Length > data1Length {
		operationLength = data2Length
	}

	result, _ := ExclusiveMerge(leftPad(data1, operationLength, 0x00), leftPad(data2, operationLength, 0x00))
	return result
}

// ExtractBytes get a byte array as subset of another byte array
func (op *littleEndianOperations) ExtractBytes(data []byte, lsbPosition, msbPosition uint64) []byte {
	var maxMsb = uint64(byteLength*len(data) - 1)

	if msbPosition <= lsbPosition || lsbPosition > maxMsb {
		return make([]byte, 0)
	}

	if msbPosition > maxMsb {
		msbPosition = maxMsb
	}

	var result = LeftShift(data, maxMsb-msbPosition)
	var correctiveShift = maxMsb - msbPosition + lsbPosition
	result = RightShift(result, correctiveShift)

	var size = computeSize(lsbPosition, msbPosition)
	return op.trim(result, size)
}

func (op *littleEndianOperations) trim(data []byte, newSize uint64) []byte {
	return data[uint64(len(data))-newSize:]
}
