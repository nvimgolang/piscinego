package main

import (
	"time"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, slovo := range s {
		z01.PrintRune(slovo)
		time.Sleep(500 * time.Millisecond) // задержка 0.5 секу
	}
}

func main() {
	PrintStr("sybau")
}
