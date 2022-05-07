package main

import (
	"fmt"
	"piscine"
	"regexp"
)

func testSplitSpace(s string) {
	result := piscine.SplitWhiteSpaces(s)
	real := regexp.MustCompile("[ \t\n]").Split(s, -1)
	fmt.Printf("%#v\n", result)
	for i := range result {
		if result[i] != real[i] {
			fmt.Printf("%#v	<- KO\n", real)
			return
		}
	}
	if len(result) != len(real) {
		fmt.Printf("%#v	<- KO\n", real)
	}
}

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))

	fmt.Println()
	testSplitSpace("Hello how are you?")
	testSplitSpace("  Hello how are you?  ")
	testSplitSpace("Hello  how  \nare \tyou?")
	testSplitSpace(" あい  うえ ")
	testSplitSpace("    \t\n")
	testSplitSpace("")
}
