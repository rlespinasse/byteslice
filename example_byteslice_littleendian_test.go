package byteslice

import "fmt"

func ExampleRSet() {
	data := []byte{0xBA, 0xCA, 0x44}
	setData := []byte{0x12, 0x11}

	fmt.Printf("%x\n", RSet(data, setData))
	// Output: bada55
}

func ExampleRUnset() {
	data := []byte{0x11, 0x11, 0x00}
	unsetData := []byte{0x01, 0x01}

	output := RUnset(data, unsetData)
	fmt.Printf("%x\n", output)
	// Output: 110100
}

func ExampleRSubset() {
	data := []byte{0xDA, 0x99, 0xBA}
	output := RSubset(data, 6, 17)
	fmt.Printf("%x\n", output)
	// Output: 0a66
}

func ExampleRToggle() {
	data := []byte{0xDA, 0x99, 0xBA}
	toogleData := []byte{0x77, 0x88, 0x11, 0xAD, 0x11, 0xAB}

	output := RToggle(data, toogleData)
	fmt.Printf("%x\n", output)
	// Output: 778811778811
}
