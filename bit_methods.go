package gobits

// ContainsBit test if a bit (position b) is contains in a
func ContainsBit(a byte, b uint8) bool {
	return ContainsBits(a, SetBit(0x00, b))
}

// SetBit change the bit (b) of a to 1
func SetBit(a byte, b uint8) byte {
	return SetBits(a, 0x01<<b)
}

// UnsetBit change the bit (b) of a to 0
func UnsetBit(a byte, b uint8) byte {
	return UnsetBits(a, 0x01<<b)
}

// GetBit get the bit from a byte
func GetBit(a byte, b uint8) byte {
	return (a >> b & 0x01) << b
}
