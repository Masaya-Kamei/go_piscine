package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for x := 2; x*x/x == x && x*x <= nb; x++ {
		if nb%x == 0 {
			return false
		}
	}
	return true
}
