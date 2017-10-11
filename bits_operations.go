package gobits

// ContainsBits test if some bits is contains in a data
func ContainsBits(data, bits byte) bool {
	return data&bits == bits
}

// SetBits affect the bits to a data
func SetBits(data, bits byte) byte {
	return data | bits
}

// UnsetBits unaffect the bits to a data
func UnsetBits(data, bits byte) byte {
	return data &^ bits
}

// ExtractBits get a byte as subset of another byte
func ExtractBits(data byte, lsbPosition, msbPosition uint8) byte {
	return data << (msb - msbPosition) >> (msb + lsbPosition - msbPosition)
}
