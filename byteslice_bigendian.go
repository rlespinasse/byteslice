package byteslice

// LUnset apply AND operation on a byte slice with an "unset" byte slice using big endian order.
func LUnset(data, unsetData []byte) []byte {
	var dataLength = len(data)
	if dataLength < 1 {
		return data
	}

	unsetDataLength := len(unsetData)
	operationLength := dataLength
	operationCut := dataLength
	if unsetDataLength > dataLength {
		operationLength = unsetDataLength
	}

	result, _ := Unset(RPad(data, operationLength, 0xFF), RPad(unsetData, operationLength, 0xFF))
	return result[:operationCut]
}

// LSet apply OR operation on a byte slice with an "set" byte slice using big endian order.
func LSet(data, setData []byte) []byte {
	dataLength := len(data)
	setDataLength := len(setData)

	operationLength := dataLength
	if setDataLength > dataLength {
		operationLength = setDataLength
	}

	result, _ := Set(RPad(data, operationLength, 0x00), RPad(setData, operationLength, 0x00))
	return result
}

// LToggle apply XOR operation on a byte slice with an "toggle" byte slice using big endian order.
func LToggle(data, toggleData []byte) []byte {
	dataLength := len(data)
	toggleDataLength := len(toggleData)

	operationLength := dataLength
	if toggleDataLength > dataLength {
		operationLength = toggleDataLength
	}

	result, _ := Toggle(RPad(data, operationLength, 0x00), RPad(toggleData, operationLength, 0x00))
	return result
}

// LSubset get the byte slice of a subset of the big endian ordered data byte defined
// by the least significant bit and the most significant bit.
func LSubset(data []byte, leastSignificantBit, mostSignificantBit uint64) []byte {
	var maxDataMostSignificantBit = uint64(maxBitsLength*len(data) - 1)

	if mostSignificantBit <= leastSignificantBit || leastSignificantBit > maxDataMostSignificantBit {
		return make([]byte, 0)
	}

	if mostSignificantBit > maxDataMostSignificantBit {
		mostSignificantBit = maxDataMostSignificantBit
	}

	var result = RShift(data, maxDataMostSignificantBit-mostSignificantBit)
	var correctiveShift = maxDataMostSignificantBit - mostSignificantBit + leastSignificantBit
	result = LShift(result, correctiveShift)

	var size = computeSize(leastSignificantBit, mostSignificantBit)
	return result[:size]
}
