package piscine

import (
	"fmt"
	"os"
)

func getStrsLen(strs []string) int {
	i := 0
	for range strs {
		i++
	}
	return i
}

func getOperationArgs() (int, rune, int, bool) {
	args := os.Args[1:]
	argc := getStrsLen(args)
	if argc != 3 {
		return 0, 0, 0, false
	}

	var v1, v2 int
	var ope rune
	_, err := fmt.Sscanf(args[0]+" "+args[1]+" "+args[2], "%d %c %d", &v1, &ope, &v2)
	if err != nil {
		return 0, 0, 0, false
	}
	return v1, ope, v2, true
}

func Doop() {
	v1, ope, v2, ok := getOperationArgs()
	if !ok {
		return
	}

	opeMap := map[rune]func(int, int) (int, bool){
		'+': opePlus, '-': opeMinus,
		'*': opeMulti, '/': opeDiv, '%': opeMod,
	}
	if _, ok := opeMap[ope]; !ok {
		return
	}

	result, ok := opeMap[ope](v1, v2)
	if !ok {
		return
	}
	fmt.Println(result)
}
