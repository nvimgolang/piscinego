package piscine

func Split(s, sep string) []string {
	// Особый случай: пустой разделитель
	if len(sep) == 0 {
		if len(s) == 0 {
			return []string{""}
		}
		result := make([]string, len(s))
		for i, char := range s {
			result[i] = string(char)
		}
		return result
	}

	// Подсчитываем количество разделителей
	count := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			count++
			i += len(sep) - 1
		}
	}

	// Создаём слайс результата
	result := make([]string, count+1)
	start := 0
	index := 0

	// Ищем разделители и извлекаем подстроки
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result[index] = s[start:i]
			index++
			start = i + len(sep)
			i += len(sep) - 1
		}
	}

	// Добавляем последнюю часть
	result[index] = s[start:]

	return result
}
