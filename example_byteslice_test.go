package byteslice

import (
	"fmt"
	"log"
)

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
	data1 := []byte{0x00, 0x00}
	data2 := []byte{0x00}
	val, err := Unset(data1, data2)
	if err == nil || val != nil {
		log.Println("Unset with two byte slices of different size needs to return an error and no value")
	}
	fmt.Printf("%x\n", val)
	// Output:
}

func ExampleSet() {
	data1 := []byte{0x00, 0x00}
	data2 := []byte{0x00}
	val, err := Set(data1, data2)
	if err == nil || val != nil {
		log.Println("Set with two byte slices of different size needs to return an error and no value")
	}
	fmt.Printf("%x\n", val)
	// Output:
}

func ExampleFlip() {
	data := []byte{0x55, 0xDA, 0xBA}
	fmt.Printf("%x\n", Flip(data))
	// Output: aa2545
}
