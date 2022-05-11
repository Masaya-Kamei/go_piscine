package piscine

import (
	"os"
)

type opCError struct{}

func (e opCError) Error() string {
	return "Invalid option c argument."
}

func atoi(s string) (int, bool) {
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
			return 0, false
		}
		if (nbr*10/10 != nbr) ||
			(sign == 1 && nbr*10+int(r-'0') < 0) ||
			(sign == -1 && nbr*10-int(r-'0') > 0) {
			return 0, false
		}
		nbr = nbr*10 + int(r-'0')*sign
	}
	return nbr, true
}

func getArgc() int {
	i := 0
	for range os.Args {
		i++
	}
	return i
}

func readOptionC() (int, error) {
	optionC := -1
	ok := false
	argc := getArgc()
	if argc <= 1 {
		return optionC, nil
	}
	if os.Args[1] == "-c" {
		if argc == 2 {
			return 0, opCError{}
		}
		optionC, ok = atoi(os.Args[2])
		if !ok || optionC <= 0 {
			return 0, opCError{}
		}
	}
	return optionC, nil
}
