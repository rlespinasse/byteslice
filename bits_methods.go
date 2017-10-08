package gobits

// ContainsBits test if b is contains in a
func ContainsBits(a, b byte) bool {
	return a&b == b
}

// SetBits affect to a, the b bits
func SetBits(a, b byte) byte {
	return a | b
}

// UnsetBits unaffect to a, the b bits
func UnsetBits(a, b byte) byte {
	return a &^ b
}
