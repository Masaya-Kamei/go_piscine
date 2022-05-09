package main

import (
	"fmt"
	"ft"
	"piscine"
)

const intSize = 32 << (^uint(0) >> 63)

func getBitString(n int) string {
	var bytes [intSize]byte
	for i := 0; i < intSize; i++ {
		bytes[63-i] = byte(((n >> i) & 1) + '0')
	}
	return string(bytes[:])
}

func main() {
	fmt.Println(piscine.ActiveBits(7))

	ft.PrintRune('\n')
	fmt.Println(getBitString(0), piscine.ActiveBits(0))
	fmt.Println(getBitString(-1), piscine.ActiveBits(-1))
	fmt.Println(getBitString(-2), piscine.ActiveBits(-2))
	intMax := 9223372036854775807
	intMin := -9223372036854775808
	fmt.Println(getBitString(intMax), piscine.ActiveBits(intMax))
	fmt.Println(getBitString(intMin), piscine.ActiveBits(intMin))
}
