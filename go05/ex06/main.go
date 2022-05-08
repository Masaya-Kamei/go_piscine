package main

import (
	"fmt"
	"piscine"
	"strings"
)

func testSplit(s, sep string) {
	result := piscine.Split(s, sep)
	real := strings.Split(s, sep)
	fmt.Printf("%#v\n", result)
	if len(result) != len(real) {
		fmt.Printf("%#v	<- KO\n", real)
		return
	}
	for i := range result {
		if result[i] != real[i] {
			fmt.Printf("%#v	<- KO\n", real)
			return
		}
	}
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))

	fmt.Println()
	testSplit("HelloHAhowHAareHAyou?", "HA")
	testSplit("HAHelloHAhowHAareHAyou?HA", "HA")
	testSplit("HelloHAHAHAHAHAhowHAareHAyou?", "HA")
	testSplit("あいうえおあいあいうえおあい", "あい")
	testSplit("HAHAHAHAHAHA", "HA")
	testSplit("", "HA")
	testSplit("あいうえお", "")
	testSplit("HA", "")
	testSplit("", "")
}
