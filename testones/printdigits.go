package main

import (
	"time"

	"github.com/01-edu/z01"
)

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
		time.Sleep(500 * time.Millisecond) // задержка 0.5 секу
	}
}
