package main

import (
	"fmt"
	"ft"
	"piscine"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"
	fmt.Printf("%#v\n", a)
	fmt.Printf("Size after compacting:%d, %#v\n", piscine.Compact(&a), a)

	ft.PrintRune('\n')
	a = make([]string, N)
	fmt.Printf("%#v\n", a)
	fmt.Printf("Size after compacting:%d, %#v\n", piscine.Compact(&a), a)
	a = make([]string, 0)
	fmt.Printf("%#v\n", a)
	fmt.Printf("Size after compacting:%d, %#v\n", piscine.Compact(&a), a)
	a = nil
	fmt.Printf("%#v\n", a)
	fmt.Printf("Size after compacting:%d, %#v\n", piscine.Compact(&a), a)
	fmt.Printf("Size after compacting:%d\n", piscine.Compact(nil))
}
