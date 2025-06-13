package main

import (
	"time"
)

func StrRev(s string) string {
	slovonaoborot := ""
	for _, pokajdoibukve := range s {
		slovonaoborot = string(pokajdoibukve) + slovonaoborot
		time.Sleep(500 * time.Millisecond) // задержка 0.5 секу
	}
	return slovonaoborot
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
