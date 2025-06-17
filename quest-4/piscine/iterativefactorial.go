package piscine

func ItertiveFactorial(nb int) int {
	if nb <= 0 || nb > 20 {
		return 0
	}

	if nb <= 1 {
		return 1
	}

	rslt := 1

	for mnojitel := 2; mnojitel <= nb; mnojitel++ {
		rslt *= mnojitel
	}
	return rslt
}
