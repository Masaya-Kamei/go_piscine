package piscine

func Unmatch(a []int) int {
	aMap := map[int]bool{}
	for _, aa := range a {
		aMap[aa] = !aMap[aa]
	}
	for _, aa := range a {
		if aMap[aa] == true {
			return aa
		}
	}
	return -1
}
