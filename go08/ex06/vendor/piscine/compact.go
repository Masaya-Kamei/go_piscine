package piscine

func countNonZeroValue(strs []string) int {
	count := 0
	for _, s := range strs {
		if s != "" {
			count++
		}
	}
	return count
}

func Compact(ptr *[]string) int {
	if ptr == nil || *ptr == nil {
		return 0
	}
	count := countNonZeroValue(*ptr)
	strs := make([]string, count)
	i := 0
	for _, s := range *ptr {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	*ptr = strs
	return count
}
