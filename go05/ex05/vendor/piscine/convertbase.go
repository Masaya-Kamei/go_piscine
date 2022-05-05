package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := atoiBase(nbr, baseFrom)
	return itoaBase(n, baseTo)
}
