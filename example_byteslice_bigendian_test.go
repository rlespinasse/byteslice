package byteslice

import "fmt"

func ExampleLSet() {
	data := []byte{0xAA, 0xCA, 0x55}
	setData := []byte{0x10, 0x12}

	fmt.Printf("%x\n", LSet(data, setData))
	// Output: bada55
}

func ExampleLToggle() {
	data := []byte{0xAB, 0xCB, 0x44}
	setData := []byte{0x11, 0x11, 0x11}

	fmt.Printf("%x\n", LToggle(data, setData))
	// Output: bada55
}

func ExampleLUnset() {
	data := []byte{0x11, 0x11, 0x10}
	unsetData := []byte{0x01, 0x01}

	output := LUnset(data, unsetData)
	fmt.Printf("%x\n", output)
	// Output: 010110
}

func ExampleLSubset() {
	data := []byte{0xDA, 0x99, 0xBA}

	output := LSubset(data, 6, 17)
	fmt.Printf("%x\n", output)
	// Output: a660
}
