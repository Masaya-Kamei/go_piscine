package piscine

func Atoi(s string) int {
	var nbr int = 0
	sign := 1

	for i, r := range s {
		if i == 0 && (r == '+' || r == '-') {
			if r == '-' {
				sign = -1
			}
			continue
		}
		if r < '0' || r > '9' {
			return 0
		}
		nbr = nbr*10 + int(r-'0')
	}
	return nbr * sign
}
