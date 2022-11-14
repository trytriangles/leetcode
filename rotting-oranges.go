func rotNeighbors(grid [][]int, row, col int) []int {
	newlyRottedPositions := []int{}
	if col > 0 && grid[row][col-1] == 1 { // west
		grid[row][col-1] = 2
		newlyRottedPositions = append(newlyRottedPositions, row, col-1)
	}
	if col+1 < len(grid[0]) && grid[row][col+1] == 1 { // east
		grid[row][col+1] = 2
		newlyRottedPositions = append(newlyRottedPositions, row, col+1)
	}
	if row+1 < len(grid) && grid[row+1][col] == 1 { // south
		grid[row+1][col] = 2
		newlyRottedPositions = append(newlyRottedPositions, row+1, col)
	}
	if row > 0 && grid[row-1][col] == 1 { // north
		grid[row-1][col] = 2
		newlyRottedPositions = append(newlyRottedPositions, row-1, col)
	}
	return newlyRottedPositions
}

func orangesRotting(grid [][]int) int {
	minutes := 0
	rotten := []int{}
	freshCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				freshCount++
			} else if grid[row][col] == 2 {
				rotten = append(rotten, row, col)
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	if len(rotten) == 0 {
		return -1
	}
	for freshCount > 0 {
		originalRottenLength := len(rotten)
		for i := 0; i < originalRottenLength; i += 2 {
			row := rotten[i]
			col := rotten[i+1]
			rotten = append(rotten, rotNeighbors(grid, row, col)...)
		}
		newlyRottenCount := len(rotten) - originalRottenLength
		freshCount -= newlyRottenCount / 2
		if newlyRottenCount == 0 && freshCount > 0 {
			return -1
		}
		minutes++
	}
	return minutes
}
