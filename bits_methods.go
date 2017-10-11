package gobits

// ContainsBits test if b is contains in a
func ContainsBits(data, bits byte) bool {
	return data&bits == bits
}

// SetBits affect to a, the b bits
func SetBits(data, bits byte) byte {
	return data | bits
}

// UnsetBits unaffect to a, the b bits
func UnsetBits(data, bits byte) byte {
	return data &^ bits
}

// ExtractBits get a byte for subset of another byte
func ExtractBits(data byte, lsb, msb uint8) byte {
	var unusedBits = 7 + msb - lsb
	return data >> msb << unusedBits >> unusedBits
}
