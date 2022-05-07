package main

import (
	"piscine"
)

func main() {
	piscine.ZTail()
}

// echo '#!/bin/zsh\ndiff <(go run . $* 2>&1) <(tail $* 2>&1)' > test.sh
