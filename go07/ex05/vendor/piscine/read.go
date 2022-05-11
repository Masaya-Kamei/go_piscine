package piscine

import "os"

type opeFuncType func(int, int) (int, error)

func getStrsLen(strs []string) int {
	i := 0
	for range strs {
		i++
	}
	return i
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

func readOpeArgs() (int, int, opeFuncType, bool) {
	args := os.Args[1:]
	if getStrsLen(args) != 3 {
		return 0, 0, nil, false
	}

	v1, ok := atoi(args[0])
	if !ok {
		return 0, 0, nil, false
	}
	v2, ok := atoi(args[2])
	if !ok {
		return 0, 0, nil, false
	}
	opeMap := map[string]opeFuncType{
		"+": opePlus, "-": opeMinus, "*": opeMulti, "/": opeDiv, "%": opeMod,
	}
	opeFunc, ok := opeMap[args[1]]
	if !ok {
		return 0, 0, nil, false
	}
	return v1, v2, opeFunc, true
}
