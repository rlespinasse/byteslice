package byteslice

const (
	maxBitsLength          = 8
	maxLeastSignificantBit = 0
	maxMostSignificantBit  = maxBitsLength - 1
)

// RBitState get the state of a specific bit of the little endian ordered data byte.
func RBitState(data byte, bit uint8) byte {
	return (data >> bit & 0x01) << bit
}

// RBitsSubset get the byte value of a subset of the little endian ordered data byte defined
// by the least significant bit and the most significant bit.
func RBitsSubset(data byte, leastSignificantBit, mostSignificantBit uint8) byte {
	return data << (maxMostSignificantBit - mostSignificantBit) >> (maxMostSignificantBit + leastSignificantBit - mostSignificantBit)
}
