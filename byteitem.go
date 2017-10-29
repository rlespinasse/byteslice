package byteslice

const (
	byteLength = 8
	lsb        = 0
	msb        = byteLength - 1
)

// ByteItem attaches the methods of Interface to byte
type ByteItem byte

// ContainsBit test if a bit is contains in the data
func (data ByteItem) ContainsBit(bit uint8) bool {
	return data.Contains(ByteItem(0x00).SetBit(bit))
}

// Contains test if some bits is contains in a data
func (data ByteItem) Contains(bits ByteItem) bool {
	return data&bits == bits
}

// SetBit change the bit of the data to 1
func (data ByteItem) SetBit(bit uint8) ByteItem {
	return data.Set(0x01 << bit)
}

// Set affect the bits to a data
func (data ByteItem) Set(bits ByteItem) ByteItem {
	return data | bits
}

// UnsetBit change the bit of the data to 0
func (data ByteItem) UnsetBit(bit uint8) ByteItem {
	return data.Unset(0x01 << bit)
}

// Unset unaffect the bits to a data
func (data ByteItem) Unset(bits ByteItem) ByteItem {
	return data &^ bits
}

// GetBit get the bit value from position
func (data ByteItem) GetBit(bit uint8) ByteItem {
	return (data >> bit & 0x01) << bit
}

// Subset get a byte as subset of another byte
func (data ByteItem) Subset(lsbPosition, msbPosition uint8) ByteItem {
	return data << (msb - msbPosition) >> (msb + lsbPosition - msbPosition)
}
