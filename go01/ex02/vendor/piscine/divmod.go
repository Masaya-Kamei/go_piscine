package piscine

func DivMod(a int, b int, div *int, mod *int) {
	if b == 0 {
		return
	}
	*div = a / b
	*mod = a % b
}
