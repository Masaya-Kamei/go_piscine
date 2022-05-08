package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)

	fmt.Println("")
	result = piscine.ConvertBase("0", "01", "0123456789")
	fmt.Println(result)
	result = piscine.ConvertBase("", "01", "0123456789")
	fmt.Println(result)
	result = piscine.ConvertBase("9223372036854775807", "0123456789", "零一二三四五六七八九")
	fmt.Println(result)
	result = piscine.ConvertBase("九二二三三七二零三六八五四七七五八零七", "零一二三四五六七八九", "0123456789")
	fmt.Println(result)
}
