package byteness

type bigEndianOperations struct{}

// BigEndianOperations test
var BigEndianOperations = &bigEndianOperations{}

// Mask apply AND mask to byte array
func (op *bigEndianOperations) Mask(data, mask []byte) []byte {
	var dataLength = len(data)
	if dataLength < 1 {
		return data
	}

	var maskLength = len(mask)
	var operationLength = dataLength
	var operationCut = dataLength
	if maskLength > dataLength {
		operationLength = maskLength
	}

	result, _ := Mask(rightPad(data, operationLength, 0xFF), rightPad(mask, operationLength, 0xFF))
	return result[:operationCut]
}

// InclusiveMerge apply OR operation on two byte arrays
func (op *bigEndianOperations) InclusiveMerge(data1, data2 []byte) []byte {
	var data1Length = len(data1)
	var data2Length = len(data2)

	var operationLength = data1Length
	if data2Length > data1Length {
		operationLength = data2Length
	}

	result, _ := InclusiveMerge(rightPad(data1, operationLength, 0x00), rightPad(data2, operationLength, 0x00))
	return result
}

// ExclusiveMerge apply XOR operation on two byte arrays
func (op *bigEndianOperations) ExclusiveMerge(data1, data2 []byte) []byte {
	var data1Length = len(data1)
	var data2Length = len(data2)

	var operationLength = data1Length
	if data2Length > data1Length {
		operationLength = data2Length
	}

	result, _ := ExclusiveMerge(rightPad(data1, operationLength, 0x00), rightPad(data2, operationLength, 0x00))
	return result
}

// ExtractBytes get a byte array as subset of another byte array
func (op *bigEndianOperations) ExtractBytes(data []byte, lsbPosition, msbPosition uint64) []byte {
	var maxMsb = uint64(byteLength*len(data) - 1)

	if msbPosition <= lsbPosition || lsbPosition > maxMsb {
		return make([]byte, 0)
	}

	if msbPosition > maxMsb {
		msbPosition = maxMsb
	}

	var result = RightShift(data, maxMsb-msbPosition)
	var correctiveShift = maxMsb - msbPosition + lsbPosition
	result = LeftShift(result, correctiveShift)

	var size = computeSize(lsbPosition, msbPosition)
	return op.trim(result, size)
}

func (op *bigEndianOperations) trim(data []byte, newSize uint64) []byte {
	return data[:newSize]
}
