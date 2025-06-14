package piscine

func SortIntegerTable(table []int) {
	ln := len(table)

	for i := 0; i < ln-1; i++ {
		for j := 0; j < ln-i-1; j++ {
			if table[j] > table[j+1] {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}
