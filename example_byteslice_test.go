package byteslice

import "fmt"

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

func ExampleRShift() {
	data := []byte{0xDA, 0x99, 0xBA}
	shift := uint64(16)
	fmt.Printf("%x\n", RShift(data, shift))
	// Output: 0000da
}

func ExampleRPad() {
	data := []byte{0xDA}
	length := 2
	filler := byte(0x00)
	fmt.Printf("%x\n", RPad(data, length, filler))
	// Output: da00
}
