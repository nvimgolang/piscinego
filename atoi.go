package piscien_test

func Atoi(s string) int {
	o_number := 0
	c := 0
	checker := true
	a_s := []rune(s)
	pl := 1

	if s != "" {
		if byte(a_s[0]) == 45 {
			pl = -1
			a_s[0] = '0'
		} else if byte(a_s[0]) == 43 {
			a_s[0] = '0'
		}
	}
}
