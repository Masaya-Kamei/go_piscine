package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	} else if nb%2 == 0 {
		return false
	}
	for x := 3; x*x/x == x && x*x <= nb; x += 2 {
		if nb%x == 0 {
			return false
		}
	}
	return true
}
