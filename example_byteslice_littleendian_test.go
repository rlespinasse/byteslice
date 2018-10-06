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
