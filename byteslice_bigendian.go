package byteslice

// BigEndianByteSlice is a ByteSlice who are big endian ordered
type BigEndianByteSlice ByteSlice

// Mask apply AND mask on the slice
func (data BigEndianByteSlice) Mask(mask BigEndianByteSlice) BigEndianByteSlice {
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

	result, _ := ByteSlice(ByteSlice(data).rightPad(operationLength, 0xFF)).Mask(ByteSlice(mask).rightPad(operationLength, 0xFF))
	return BigEndianByteSlice(result).newSize(operationCut)
}

func (data BigEndianByteSlice) newSize(newSize int) BigEndianByteSlice {
	return []byte(data)[:newSize]
}

// InclusiveMerge apply OR operation between this slice and another
func (data BigEndianByteSlice) InclusiveMerge(anotherData BigEndianByteSlice) BigEndianByteSlice {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)

	var operationLength = dataLength
	if anotherDataLength > dataLength {
		operationLength = anotherDataLength
	}

	result, _ := ByteSlice(ByteSlice(data).rightPad(operationLength, 0x00)).InclusiveMerge(ByteSlice(anotherData).rightPad(operationLength, 0x00))
	return BigEndianByteSlice(result)
}

// ExclusiveMerge apply XOR operation between this slice and another
func (data BigEndianByteSlice) ExclusiveMerge(anotherData BigEndianByteSlice) BigEndianByteSlice {
	var dataLength = len(data)
	var anotherDataLength = len(anotherData)

	var operationLength = dataLength
	if anotherDataLength > dataLength {
		operationLength = anotherDataLength
	}

	result, _ := ByteSlice(ByteSlice(data).rightPad(operationLength, 0x00)).ExclusiveMerge(ByteSlice(anotherData).rightPad(operationLength, 0x00))

	return BigEndianByteSlice(result)
}

// Subset extract a subset of this slice
func (data BigEndianByteSlice) Subset(lsbPosition, msbPosition uint64) BigEndianByteSlice {
	var maxMsb = uint64(byteLength*len(data) - 1)

	if msbPosition <= lsbPosition || lsbPosition > maxMsb {
		return make(BigEndianByteSlice, 0)
	}

	if msbPosition > maxMsb {
		msbPosition = maxMsb
	}

	var result = ByteSlice(data).RightShift(maxMsb - msbPosition)
	var correctiveShift = maxMsb - msbPosition + lsbPosition
	result = ByteSlice(result).LeftShift(correctiveShift)

	var size = computeSize(lsbPosition, msbPosition)
	return BigEndianByteSlice(result).trim(size)
}

// LittleEndianSubset get a byte array as subset of another byte array when byte are in little endian
func (data BigEndianByteSlice) LittleEndianSubset(lsbPosition, msbPosition uint64) BigEndianByteSlice {
	var reversedData = ByteSlice(data).Reverse()
	var result = LittleEndianByteSlice(reversedData).Subset(lsbPosition, msbPosition)
	return BigEndianByteSlice(ByteSlice(result).Reverse())
}

func (data BigEndianByteSlice) trim(newSize uint64) BigEndianByteSlice {
	return data[:newSize]
}
