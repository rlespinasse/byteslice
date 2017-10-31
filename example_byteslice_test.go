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

func ExampleUnset() {
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	zeroes := []byte{0x00, 0x00, 0x00, 0x00}

	dat, err := Unset(data, zeroes)
	fmt.Println(err, dat)
	// Output: <nil> [0 0 0 0]
}
