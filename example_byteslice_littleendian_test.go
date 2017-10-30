package byteslice

import "fmt"

func ExampleRSet() {
	data := []byte{0xBA, 0xCA, 0x44}
	setData := []byte{0x12, 0x11}

	fmt.Printf("%x\n", RSet(data, setData))
	// Output: bada55
}

func ExampleRUnset() {
	data := []byte{0xDA, 0x99, 0xBA}
	unsetData := []byte{0xAD, 0x11, 0xAB}
	fmt.Printf("%x\n", RUnset(data, unsetData))
	// Output: 8811aa
}

func ExampleRToogle() {
	data := []byte{0xDA, 0x99, 0xBA}
	toogleData := []byte{0xAD, 0x11, 0xAB}
	fmt.Printf("%x\n", RToogle(data, toogleData))
	// Output: 778811
}

func ExampleRSubset() {
	data := []byte{0xDA, 0x99, 0xBA}
	leastSignificantBit := (uint64)(100)
	mostSignificantBit := (uint64)(101)
	fmt.Printf("%x\n", RSubset(data, leastSignificantBit, mostSignificantBit))
	// Output:
}
