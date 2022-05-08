package piscine

import "fmt"

func opePlus(v1, v2 int) (int, bool) {
	result := v1 + v2
	if v1 > 0 && v2 > 0 && result <= 0 ||
		v1 < 0 && v2 < 0 && result >= 0 {
		return 0, false
	}
	return result, true
}

func opeMinus(v1, v2 int) (int, bool) {
	result := v1 - v2
	if v1 == 0 && v2 == -v2 ||
		v1 > 0 && v2 < 0 && result <= 0 ||
		v1 < 0 && v2 > 0 && result >= 0 {
		return 0, false
	}
	return result, true
}

func opeMulti(v1, v2 int) (int, bool) {
	result := v1 * v2
	if result != 0 && result/v1 != v2 {
		return 0, false
	}
	return result, true
}

func opeDiv(v1, v2 int) (int, bool) {
	if v2 == 0 {
		fmt.Println("No division by 0")
		return 0, false
	}
	return v1 / v2, true
}

func opeMod(v1, v2 int) (int, bool) {
	if v2 == 0 {
		fmt.Println("No modulo by 0")
		return 0, false
	}
	return v1 % v2, true
}
