package piscine

func nilCheck2(p **int) bool {
	return p == nil || *p == nil
}

func nilCheck3(p ***int) bool {
	return p == nil || nilCheck2(*p)
}

func nilCheck4(p ****int) bool {
	return p == nil || nilCheck3(*p)
}

func Enigma(a ***int, b *int, c *******int, d ****int) {
	if nilCheck3(a) || b == nil || nilCheck4(d) ||
		c == nil || *c == nil || **c == nil || nilCheck4(***c) {
		return
	}
	***a, *b, *******c, ****d = *b, ****d, ***a, *******c
}
