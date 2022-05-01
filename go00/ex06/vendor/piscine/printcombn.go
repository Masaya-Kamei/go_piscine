package piscine

import "ft"

func printNumber(n int, num [9]rune) {
	for i := 0; i < n; i++ {
		ft.PrintRune(num[i])
	}
}

func getEndRune(n, i int) rune {
	return rune('9' - n + 1 + i)
}

func printCombNRec(n int, num [9]rune, i int) {
	end := getEndRune(n, i)

	if i == 0 {
		num[0] = '0'
	} else {
		num[i] = num[i-1] + 1
	}
	for ; num[i] <= end; num[i]++ {
		if n != i+1 {
			printCombNRec(n, num, i+1)
		} else {
			printNumber(n, num)
			if num[0] == getEndRune(n, 0) {
				break
			}
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
	}
}

func PrintCombN(n int) {
	var num [9]rune

	if n < 1 || n > 9 {
		return
	}
	printCombNRec(n, num, 0)
	ft.PrintRune('\n')
}
