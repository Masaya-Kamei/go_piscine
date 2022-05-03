package piscine

func Join(strs []string, sep string) string {
	join := ""
	for i, s := range strs {
		if i != 0 {
			join += sep
		}
		join += s
	}
	return join
}
