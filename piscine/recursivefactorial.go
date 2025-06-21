package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb <= 1 {
		return 1
	}
	if nb > 20 {
		return 0
	}
	return nb * RecursiveFactorial(nb-1)
}
