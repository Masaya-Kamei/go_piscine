package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))

	// s = "HAHelloHAhowHAareHAyou?HA"
	// fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	// s = "HelloHAHAHAHAHAhowHAareHAyou?"
	// fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	// s = "あいうえおあいあいうえおあい"
	// fmt.Printf("%#v\n", piscine.Split(s, "あい"))
	// s = "HAHAHAHAHAHA"
	// fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
