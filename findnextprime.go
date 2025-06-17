package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	for !isPrime(nb) {
		nb++
	}
	return nb
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Check odd divisors up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
