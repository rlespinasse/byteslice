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

func ExampleToogle() {
	data1 := []byte{0x55, 0xDA, 0xBA}
	data2 := []byte{0x01, 0x23, 0x45}

	dat, err := Toogle(data1, data2)
	fmt.Printf("%v %x\n", err, dat)
	// Output: <nil> 54f9ff
}
