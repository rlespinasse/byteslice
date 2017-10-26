package byteness

const (
	byteLength = 8
	lsb        = 0
	msb        = byteLength - 1
)

// ContainsBit test if a bit is contains in the data
func ContainsBit(data byte, bit uint8) bool {
	return ContainsBits(data, SetBit(0x00, bit))
}

// SetBit change the bit of the data to 1
func SetBit(data byte, bit uint8) byte {
	return SetBits(data, 0x01<<bit)
}

// UnsetBit change the bit of the data to 0
func UnsetBit(data byte, bit uint8) byte {
	return UnsetBits(data, 0x01<<bit)
}

// GetBit get the bit value from the data
func GetBit(data byte, bit uint8) byte {
	return (data >> bit & 0x01) << bit
}
