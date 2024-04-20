package retrier

import "fmt"

func ExampleParseHexString() {
	out1, _ := ParseHexString("010203AABBCCDD")
	fmt.Println(out1)

	out2, _ := ParseHexString("0xFF")
	fmt.Println(out2)

	out3, _ := ParseHexString("invalid_hex_string")
	fmt.Println(out3)

	// Output:
	// [1 2 3 170 187 204 221]
	// [255]
	// []
}
