package main

import "fmt"

func BasicAtoi(s string) int {
	o_number := 0
	c := 0
	a_s := []rune(s)
	for _, word := range a_s {
		for i := '0'; i < word; i++ {
			c++
		}
		o_number = o_number*10 + c
		c = 0
	}
	return o_number
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
