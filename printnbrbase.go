package piscine

import "github.com/01-edu/z01"

// PrintNbrBase выводит число nbr в системе счисления, заданной строкой base
// Если base невалидна, выводит "NV" (Not Valid)
func PrintNbrBase(nbr int, base string) {
	// Проверяем валидность базы
	if !isValidBase(base) {
		// Если база невалидна, выводим "NV"
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Получаем длину базы (основание системы счисления)
	baseLen := len(base)

	// Обрабатываем отрицательные числа
	if nbr < 0 {
		// Выводим знак минус
		z01.PrintRune('-')
		// Работаем с абсолютным значением
		// Используем рекурсию для положительного числа
		printPositiveNbrBase(-nbr, base, baseLen)
	} else {
		// Для положительных чисел вызываем вспомогательную функцию
		printPositiveNbrBase(nbr, base, baseLen)
	}
}

// printPositiveNbrBase рекурсивно выводит положительное число в заданной базе
func printPositiveNbrBase(nbr int, base string, baseLen int) {
	// Базовый случай рекурсии: если число меньше основания базы
	if nbr < baseLen {
		// Выводим соответствующий символ из базы
		z01.PrintRune(rune(base[nbr]))
		return
	}

	// Рекурсивный случай: сначала выводим старшие разряды
	printPositiveNbrBase(nbr/baseLen, base, baseLen)
	// Затем выводим текущий младший разряд
	z01.PrintRune(rune(base[nbr%baseLen]))
}

// isValidBase проверяет валидность базы согласно правилам:
// 1. База должна содержать минимум 2 символа
// 2. Каждый символ должен быть уникальным
// 3. База не должна содержать символы '+' или '-'
func isValidBase(base string) bool {
	// Правило 1: база должна содержать минимум 2 символа
	if len(base) < 2 {
		return false
	}

	// Проверяем каждый символ в базе
	for i, char := range base {
		// Правило 3: база не должна содержать '+' или '-'
		if char == '+' || char == '-' {
			return false
		}

		// Правило 2: проверяем уникальность символа
		// Сравниваем текущий символ со всеми последующими
		for j := i + 1; j < len(base); j++ {
			if char == rune(base[j]) {
				return false // Найден дублирующийся символ
			}
		}
	}

	// Если все проверки пройдены, база валидна
	return true
}
