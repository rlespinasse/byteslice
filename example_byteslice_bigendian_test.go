package byteslice

import "fmt"

func ExampleLSet() {
	data := []byte{0xAA, 0xCA, 0x55}
	setData := []byte{0x10, 0x12}

	fmt.Printf("%x\n", LSet(data, setData))
	// Output: bada55
}

func ExampleLUnset() {
	data := []byte{0xBA, 0xDA, 0x55}
	setData := []byte{0xBA, 0xDA}

  fmt.Printf("%x\n", LUnset(data, setData))
	// Output: bada55
}

func ExampleLToggle() {
	data := []byte{0xAB, 0xCB, 0x44}
	setData := []byte{0x11, 0x11, 0x11}

	fmt.Printf("%x\n", LToggle(data, setData))
	// Output: bada55
}
