package piscine

// QuadA рисует прямоугольник с углами 'o'
func QuadA(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			isCorner := (row == 1 || row == height) &&
				(col == 1 || col == width)
			isTopBottom := row == 1 || row == height
			isLeftRight := col == 1 || col == width

			switch {
			case isCorner:
				print("o")
			case isTopBottom:
				print("-")
			case isLeftRight:
				print("|")
			default:
				print(" ")
			}
		}
		print("\n")
	}
}

