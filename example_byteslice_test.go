package byteslice

import "fmt"

func ExampleReverse() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", Reverse(data))
	// Output: bada55
}

func ExampleRPad() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", RPad(data, 5, 0x22))
	// Output: 55daba2222
}

func ExampleLPad() {
	data := []byte{0x55, 0xDA, 0xBA}

	fmt.Printf("%x\n", LPad(data, 5, 0x22))
	// Output: 222255daba
}

func ExampleToggle() {
	data := []byte{0xbb, 0xdb, 0x54}
	toggle := []byte{0x01, 0x01, 0x01}

	output, _ := Toggle(data, toggle)
	fmt.Printf("%x\n", output)
	// Output: bada55
}

func ExampleToggle_simple() {
	data := []byte{0x00, 0x01, 0x00}
	toggle := []byte{0x01, 0x00, 0x01}

	output, _ := Toggle(data, toggle)
	fmt.Printf("%x\n", output)
	// Output: 010101
}

func ExampleUnset() {
	data := []byte{0x01, 0x11, 0x00}
	unsetData := []byte{0x01, 0x01, 0x01}

	output, _ := Unset(data, unsetData)
	fmt.Printf("%x\n", output)
	// Output: 010100
}

func ExampleLShift() {
	data := []byte{0xDA, 0x99, 0xBA}

	output := LShift(data, 8)
	fmt.Printf("%x\n", output)
	// Output: 99ba00
}

func ExampleRShift() {
	data := []byte{0xDA, 0x99, 0xBA}

	output := RShift(data, 8)
	fmt.Printf("%x\n", output)
	// Output: 00da99
}

func ExampleSet() {
	data := []byte{0xDA, 0x99, 0xBA}
	setData := []byte{0x11, 0x22, 0x33}

	output, _ := Set(data, setData)
	fmt.Printf("%x\n", output)
	// Output: dbbbbb
}

func ExampleFlip() {
	data := []byte{0xDA, 0x99, 0xBA}

	output := Flip(data)
	fmt.Printf("%x\n", output)
	// Output: 256645
}
