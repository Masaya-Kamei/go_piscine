package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := atoiBase(nbr, baseFrom)
	if n == -1 {
		return "NV"
	}
	return itoaBase(n, baseTo)
}
