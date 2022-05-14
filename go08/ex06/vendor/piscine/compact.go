package piscine

func getNoEmptyStrMap(strs []string) (map[int]string, int) {
	mapLen := 0
	strMap := map[int]string{}
	for _, s := range strs {
		if s != "" {
			strMap[mapLen] = s
			mapLen++
		}
	}
	return strMap, mapLen
}

func Compact(ptr *[]string) int {
	if ptr == nil || *ptr == nil {
		return 0
	}
	strMap, mapLen := getNoEmptyStrMap(*ptr)
	*ptr = make([]string, mapLen)
	for i := 0; i < mapLen; i++ {
		(*ptr)[i] = strMap[i]
	}
	return mapLen
}
