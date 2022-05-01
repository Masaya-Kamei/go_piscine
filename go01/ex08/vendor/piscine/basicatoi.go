package piscine

func BasicAtoi(s string) int {
	var nbr int
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		nbr = nbr*10 + int(r-'0')
	}
	return nbr
}
