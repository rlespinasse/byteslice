package byteslice

import "fmt"

func ExampleRSet() {
	data := []byte{0xBA, 0xCA, 0x44}
	setData := []byte{0x12, 0x11}

	fmt.Printf("%x\n", RSet(data, setData))
	// Output: bada55
}
