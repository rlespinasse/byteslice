package byteslice

import "fmt"

func ExampleRBit() {
	fmt.Printf("%x\n", RBit(0x55, 6))
	// Output: 40
}

func ExampleRBitsSubset() {
	fmt.Printf("%x\n", RBitsSubset(0x55, 2, 6))
	// Output: 15
}
