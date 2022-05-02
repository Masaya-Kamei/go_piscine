package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb <= 1 {
		return 1
	}
	ret := RecursiveFactorial(nb - 1)
	if ret*nb/nb != ret {
		return 0
	}
	return ret * nb
}
