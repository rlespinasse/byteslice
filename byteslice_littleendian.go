package byteslice

// RUnset apply AND operation on a byte slice with an "unset" byte slice using little endian order.
func RUnset(data, unsetData []byte) []byte {
	var dataLength = len(data)
	if dataLength < 1 {
		return data
	}

	unsetDataLength := len(unsetData)
	operationLength := dataLength
	operationCut := 0
	if unsetDataLength > dataLength {
		operationLength = unsetDataLength
		operationCut = operationLength - dataLength
	}

	result, _ := Unset(LPad(data, operationLength, 0xFF), LPad(unsetData, operationLength, 0xFF))
	return result[operationCut:]
}

// RSet apply OR operation on a byte slice with an "set" byte slice using little endian order.
func RSet(data, setData []byte) []byte {
	dataLength := len(data)
	setDataLength := len(setData)

	operationLength := dataLength
	if setDataLength > dataLength {
		operationLength = setDataLength
	}

	result, _ := Set(LPad(data, operationLength, 0x00), LPad(setData, operationLength, 0x00))
	return result
}

// RToogle apply XOR operation on a byte slice with an "toogle" byte slice using little endian order.
func RToogle(data, toogleData []byte) []byte {
	dataLength := len(data)
	toogleDataLength := len(toogleData)

	operationLength := dataLength
	if toogleDataLength > dataLength {
		operationLength = toogleDataLength
	}

	result, _ := Toogle(LPad(data, operationLength, 0x00), LPad(toogleData, operationLength, 0x00))
	return result
}

// RSubset get the byte slice of a subset of the little endian ordered data byte defined
// by the least significant bit and the most significant bit.
func RSubset(data []byte, leastSignificantBit, mostSignificantBit uint64) []byte {
	var maxDataMostSignificantBit = uint64(maxBitsLength*len(data) - 1)

	if mostSignificantBit <= leastSignificantBit || leastSignificantBit > maxDataMostSignificantBit {
		return make([]byte, 0)
	}

	if mostSignificantBit > maxDataMostSignificantBit {
		mostSignificantBit = maxDataMostSignificantBit
	}

	var result = LShift(data, maxDataMostSignificantBit-mostSignificantBit)
	var correctiveShift = maxDataMostSignificantBit - mostSignificantBit + leastSignificantBit
	result = RShift(result, correctiveShift)

	var size = computeSize(leastSignificantBit, mostSignificantBit)
	return result[uint64(len(result))-size:]
}
