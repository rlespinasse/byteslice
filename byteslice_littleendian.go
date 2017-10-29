package byteslice

// LittleEndianByteSlice is a ByteSlice who are little endian ordered
type LittleEndianByteSlice ByteSlice

// Mask apply AND mask to a byte array
func (data LittleEndianByteSlice) Mask(mask LittleEndianByteSlice) LittleEndianByteSlice {
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

	result, _ := ByteSlice(ByteSlice(data).leftPad(operationLength, 0xFF)).Mask(ByteSlice(mask).leftPad(operationLength, 0xFF))
	return LittleEndianByteSlice(result).newSize(operationCut)
}

func (data LittleEndianByteSlice) newSize(newSize int) LittleEndianByteSlice {
	return []byte(data)[newSize:]
}

// InclusiveMerge apply OR operation on two byte arrays
func (data LittleEndianByteSlice) InclusiveMerge(anotherData LittleEndianByteSlice) LittleEndianByteSlice {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)

	var operationLength = dataLength
	if anotherDataLength > dataLength {
		operationLength = anotherDataLength
	}

	result, _ := ByteSlice(ByteSlice(data).leftPad(operationLength, 0x00)).InclusiveMerge(ByteSlice(anotherData).leftPad(operationLength, 0x00))
	return LittleEndianByteSlice(result)
}

// ExclusiveMerge apply XOR operation on two byte arrays
func (data LittleEndianByteSlice) ExclusiveMerge(anotherData LittleEndianByteSlice) LittleEndianByteSlice {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)

	var operationLength = dataLength
	if anotherDataLength > dataLength {
		operationLength = anotherDataLength
	}

	result, _ := ByteSlice(ByteSlice(data).leftPad(operationLength, 0x00)).ExclusiveMerge(ByteSlice(anotherData).leftPad(operationLength, 0x00))
	return LittleEndianByteSlice(result)
}

// Subset get a byte array as subset of another byte array
func (data LittleEndianByteSlice) Subset(lsbPosition, msbPosition uint64) LittleEndianByteSlice {
	var maxMsb = uint64(byteLength*len(data) - 1)

	if msbPosition <= lsbPosition || lsbPosition > maxMsb {
		return make(LittleEndianByteSlice, 0)
	}

	if msbPosition > maxMsb {
		msbPosition = maxMsb
	}

	var result = ByteSlice(data).LeftShift(maxMsb - msbPosition)
	var correctiveShift = maxMsb - msbPosition + lsbPosition
	result = ByteSlice(result).RightShift(correctiveShift)

	var size = computeSize(lsbPosition, msbPosition)
	return LittleEndianByteSlice(result).trim(size)
}

// BigEndianSubset get a byte array as subset of another byte array when byte are in big endian
func (data LittleEndianByteSlice) BigEndianSubset(lsbPosition, msbPosition uint64) LittleEndianByteSlice {
	var reversedData = ByteSlice(data).Reverse()
	var result = BigEndianByteSlice(reversedData).Subset(lsbPosition, msbPosition)
	return LittleEndianByteSlice(ByteSlice(result).Reverse())
}

func (data LittleEndianByteSlice) trim(newSize uint64) LittleEndianByteSlice {
	return data[uint64(len(data))-newSize:]
}
