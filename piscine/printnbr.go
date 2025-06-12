package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	hashmap := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}

	digits := []rune{}
	for n > 0 {
		digits = append(digits, hashmap[n%10])
		n = n / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
