package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	if a == nil || *a == nil || **a == nil ||
		b == nil ||
		c == nil || *c == nil || **c == nil || ***c == nil ||
		****c == nil || *****c == nil || ******c == nil ||
		d == nil || *d == nil || **d == nil || ***d == nil {
		return
	}
	***a, *b, *******c, ****d = *b, ****d, ***a, *******c
}
