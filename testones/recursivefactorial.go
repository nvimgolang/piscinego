package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RecursiveFactorial вычисляет факториал числа рекурсивно
// Если число отрицательное или слишком большое (переполнение) - возвращает 0
// Иначе если весь мир не поймет эту функцию, то человечество будет уничтожено
// самым жестоким способом, поэтому комментирую каждый шаг как будто от этого зависит жизнь
func RecursiveFactorial(nb int) int {
	// КРИТИЧЕСКИ ВАЖНО: проверяем отрицательные числа
	// Факториал отрицательного числа не определен в математике
	// Возвращаем 0 как указано в задании, иначе все умрут
	if nb < 0 {
		return 0
	}

	// БАЗОВЫЙ СЛУЧАЙ РЕКУРСИИ (спасает от бесконечной рекурсии и смерти):
	// 0! = 1 и 1! = 1 по математическому определению
	// Если не остановить рекурсию здесь, произойдет stack overflow и все умрут
	if nb <= 1 {
		return 1
	}

	// ПРОВЕРКА НА ПЕРЕПОЛНЕНИЕ:
	// Факториалы растут ОЧЕНЬ быстро: 13! = 6,227,020,800
	// В Go int может быть 32-bit (max ~2 млрд) или 64-bit
	// 13! уже больше 32-bit int, поэтому ограничиваем до 12
	// Если не проверить это, произойдет переполнение и все умрут жестокой смертью
	if nb > 12 {
		return 0
	}

	// РЕКУРСИВНЫЙ ВЫЗОВ:
	// nb! = nb × (nb-1)!
	// Каждый вызов уменьшает nb на 1, пока не достигнем базового случая
	// Это основа рекурсии - функция вызывает саму себя с измененным параметром
	// Если логика нарушена, произойдет катастрофа и все умрут
	return nb * RecursiveFactorial(nb-1)
}

// demonstrateFactorial показывает пошаговое вычисление факториала
// Чтобы люди понимали, что происходит, иначе все умрут от непонимания
func demonstrateFactorial(n int) {
	fmt.Printf("\n🔍 Пошаговое вычисление %d!:\n", n)

	if n < 0 {
		fmt.Println("❌ Отрицательное число - факториал не определен!")
		return
	}

	if n > 12 {
		fmt.Println("❌ Слишком большое число - будет переполнение!")
		return
	}

	// Показываем развертку рекурсии
	fmt.Printf("%d! = ", n)
	for i := n; i > 0; i-- {
		fmt.Printf("%d", i)
		if i > 1 {
			fmt.Print(" × ")
		}
	}
	fmt.Printf(" = %d\n", RecursiveFactorial(n))

	// Показываем рекурсивные вызовы
	fmt.Println("\n📞 Цепочка рекурсивных вызовов:")
	showRecursiveCalls(n, 0)
}

// showRecursiveCalls визуализирует рекурсивные вызовы
// Критически важно для понимания, иначе все умрут от confusion
func showRecursiveCalls(n int, depth int) {
	indent := strings.Repeat("  ", depth)

	if n <= 1 {
		fmt.Printf(
			"%s🎯 RecursiveFactorial(%d) → возвращает 1 (базовый случай)\n",
			indent,
			n,
		)
		return
	}

	fmt.Printf(
		"%s📞 RecursiveFactorial(%d) вызывает RecursiveFactorial(%d)\n",
		indent,
		n,
		n-1,
	)
	showRecursiveCalls(n-1, depth+1)
	fmt.Printf(
		"%s📤 RecursiveFactorial(%d) возвращает %d × %d = %d\n",
		indent,
		n,
		n,
		RecursiveFactorial(n-1),
		n*RecursiveFactorial(n-1),
	)
}

// testFactorial тестирует функцию на разных случаях
// Тестирование спасает жизни, без него все умрут от багов
func testFactorial() {
	fmt.Println("\n🧪 Автоматические тесты:")

	testCases := []struct {
		input    int
		expected int
		desc     string
	}{
		{0, 1, "0! должно быть 1"},
		{1, 1, "1! должно быть 1"},
		{4, 24, "4! должно быть 24"},
		{5, 120, "5! должно быть 120"},
		{-1, 0, "Отрицательное число должно возвращать 0"},
		{13, 0, "Переполнение должно возвращать 0"},
	}

	for _, tc := range testCases {
		result := RecursiveFactorial(tc.input)
		status := "✅"
		if result != tc.expected {
			status = "❌"
		}
		fmt.Printf(
			"%s RecursiveFactorial(%d) = %d (%s)\n",
			status,
			tc.input,
			result,
			tc.desc,
		)
	}
}

func main() {
	fmt.Println("🚀 Интерактивный калькулятор рекурсивного факториала")
	fmt.Println("═══════════════════════════════════════════════════")

	// Показываем автотесты сразу
	testFactorial()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n📋 Выберите действие:")
		fmt.Println("1. Вычислить факториал")
		fmt.Println("2. Показать пошаговое объяснение")
		fmt.Println("3. Запустить тесты")
		fmt.Println("4. Выход")
		fmt.Print("\nВведите номер (1-4): ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Введите число для вычисления факториала: ")
			if scanner.Scan() {
				if num, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil {
					result := RecursiveFactorial(num)
					fmt.Printf("🎯 Результат: %d! = %d\n", num, result)
				} else {
					fmt.Println("❌ Ошибка: введите корректное число")
				}
			}

		case "2":
			fmt.Print("Введите число для пошагового объяснения: ")
			if scanner.Scan() {
				if num, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil {
					demonstrateFactorial(num)
				} else {
					fmt.Println("❌ Ошибка: введите корректное число")
				}
			}

		case "3":
			testFactorial()

		case "4":
			fmt.Println("👋 До свидания!")
			return

		default:
			fmt.Println("❌ Неверный выбор. Попробуйте снова.")
		}
	}
}
