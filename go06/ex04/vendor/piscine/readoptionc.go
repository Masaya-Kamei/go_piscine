package piscine

import (
	"os"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

type StrError string

func (e StrError) Error() string {
	return string(e)
}

func atoi(s string) int {
	sign := 1
	if s != "" && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	nbr := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		if nbr*10/10 != nbr {
			return 0
		}
		nbr *= 10
		if n := nbr + int(r-'0'); n < 0 && !(sign == -1 && n == MinInt) {
			return 0
		}
		nbr += int(r - '0')
	}
	return nbr * sign
}

func isExistElement(strs []string, index int) bool {
	for i := range strs {
		if i == index {
			return true
		}
	}
	return false
}

func readOptionC() (int, error) {
	optionC := -1
	if !isExistElement(os.Args, 1) {
		return optionC, nil
	}
	if os.Args[1] == "-c" {
		if !isExistElement(os.Args, 2) {
			return 0, StrError("Invalid option c argument.")
		}
		optionC = atoi(os.Args[2])
		if optionC <= 0 {
			return 0, StrError("Invalid option c argument.")
		}
	}
	return optionC, nil
}
