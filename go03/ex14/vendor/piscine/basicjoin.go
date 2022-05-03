package piscine

func BasicJoin(elems []string) string {
	join := ""
	for _, s := range elems {
		join += s
	}
	return join
}
