package piscine

func QuadC(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	for row := 1; row <= height; row++ {
		for col := 1; col <= width; col++ {
			isTopCorner := row == 1 && (col == 1 || col == width)
			isBottomCorner := row == height && (col == 1 || col == width)
			isAnyCorner := isTopCorner || isBottomCorner
			isEdge := (row == 1 || row == height || col == 1 || col == width) &&
				!isAnyCorner

			switch {
			case isTopCorner:
				print("A")
			case isBottomCorner:
				print("C")
			case isEdge:
				print("B")
			default:
				print(" ")
			}
		}
		print("\n")
	}
}
