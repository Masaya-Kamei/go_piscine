package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	prev1 := Fibonacci(index - 1)
	prev2 := Fibonacci(index - 2)
	if prev1 < 0 || prev2 < 0 || prev1+prev2-prev2 != prev1 {
		return -1
	}
	return prev1 + prev2
}
