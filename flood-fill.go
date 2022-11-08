// Inlined by the compiler.
func isValidPosition(position []int, width, height int) bool {
	if position[0] < 0 {
		return false
	}
	if position[0] >= height {
		return false
	}
	if position[1] < 0 {
		return false
	}
	if position[1] >= width {
		return false
	}
	return true
}

// Inlined by the compiler.
func getDefaultNeighbors(row, col int) [][]int {
	return [][]int{
		{row - 1, col},
		{row, col + 1},
		{row + 1, col},
		{row, col - 1},
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	myColor := image[sr][sc]
	if myColor == color {
		return image
	}
	image[sr][sc] = color
	width := len(image[0])
	height := len(image)
	neighborPositions := getDefaultNeighbors(sr, sc)
	for _, position := range neighborPositions {
		row := position[0]
		col := position[1]
		if !isValidPosition(position, width, height) {
			continue
		}
		if image[row][col] == myColor {
			floodFill(image, row, col, color)
		}
	}
	return image
}

