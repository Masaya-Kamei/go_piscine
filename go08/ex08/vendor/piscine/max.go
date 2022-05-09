package piscine

const (
	intMax = int(^uint(0) >> 1)
	intMin = -intMax - 1
)

func isEmpty(a []int) bool {
	for range a {
		return false
	}
	return true
}

func Max(a []int) int {
	if isEmpty(a) {
		return 0
	}
	max := a[0]
	for _, aa := range a[1:] {
		if max < aa {
			max = aa
		}
	}
	return max
}
