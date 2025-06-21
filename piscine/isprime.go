package piscine

func IsPrime(nb int) bool {
	// Простые числа должны быть больше 1
	if nb <= 1 {
		return false
	}

	// 2 - единственное четное простое число
	if nb == 2 {
		return true
	}

	// Все остальные четные числа не простые
	if nb%2 == 0 {
		return false
	}

	// Проверяем деление только на нечетные числа до квадратного корня
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
