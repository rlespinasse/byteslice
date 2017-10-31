package byteslice

import "fmt"

func ExampleFlip() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", Flip(data))
	// Output: aa2545
}

func ExampleReverse() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", Reverse(data))
	// Output: bada55
}

func ExampleLPad() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", LPad(data, 5, 0x22))
	// Output: 222255daba
}
