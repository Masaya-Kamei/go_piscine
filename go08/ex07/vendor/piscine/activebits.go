package piscine

const intSize = 32 << (^uint(0) >> 63)

func ActiveBits(n int) int {
	count := 0
	for i := 0; i < intSize; i++ {
		if (n>>i)&1 == 1 {
			count++
		}
	}
	return count
}
