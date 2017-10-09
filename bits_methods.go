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

// ExtractBits get a byte for subset of another byte
func ExtractBits(a byte, b, c uint8) byte {
	var unusedBits = 7 + b - c
	return a >> b << unusedBits >> unusedBits
}
