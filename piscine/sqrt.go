package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}

	// Ищем квадратный корень от 1 до nb
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
		if i*i > nb {
			return 0
		}
	}
	return 0
}
