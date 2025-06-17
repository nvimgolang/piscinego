package piscine

func IterativePower(chislo int, stepen int) int {
	// Если степень отрицательная, возвращаем 0
	if stepen < 0 {
		return 0
	}

	// Если степень равна 0, любое число в степени 0 равно 1
	if stepen == 0 {
		return 1
	}

	// Инициализируем результат как 1
	rezultat := 1

	// Итеративно умножаем результат на число степень раз
	for schetchik := 0; schetchik < stepen; schetchik++ {
		rezultat = rezultat * chislo
	}

	return rezultat
}
