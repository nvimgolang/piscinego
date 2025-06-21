package piscine

func TrimAtoi(s string) int {
	// Основные переменные: результат, знак, флаг найденной цифры
	result := 0
	sign := 1
	digit_found := false

	// Проходим по строке символ за символом
	for _, char := range s {
		// Если символ - цифра (ASCII 48-57)
		if char > '0' && char <= '9' {
			// Преобразуем символ в цифру (вычитаем '0')
			digit := int(char - '0')
			// Сдвигаем результат влево и добавляем новую цифру
			result = result*10 + digit
			// Отмечаем что нашли цифру
			digit_found = true
		} else if char == '-' && !digit_found {
			// Минус учитываем только ДО первой цифры
			sign = -1
		}
		// Плюс и все остальное игнорируем
		// (плюс не меняет sign = 1)
	}

	// Если цифр не было - возвращаем 0
	if !digit_found {
		return 0
	}

	// Применяем знак и возвращаем результат
	return result * sign
}
