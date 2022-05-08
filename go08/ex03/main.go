package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	x := 5
	y := &x
	z := &y
	a := &z
	w := 2
	b := &w
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	ft.PrintRune('\n')
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	piscine.Enigma(a, b, c, nil)
	piscine.Enigma(a, b, nil, d)
	piscine.Enigma(a, nil, c, d)
	piscine.Enigma(nil, b, c, d)
}
