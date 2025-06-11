package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for i2 := '0'; i2 <= '9'; i2++ {
				for j2 := '0'; j2 <= '9'; j2++ {
					if i == i2 && j < j2 {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(i2)
						z01.PrintRune(j2)
						if i != '9' || j != '8' || i2 != '9' || j2 != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					} else if i < i2 {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(i2)
						z01.PrintRune(j2)
						if i != '9' || j != '8' || i2 != '9' || j2 != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
