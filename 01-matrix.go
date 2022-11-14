func updateMatrix(mat [][]int) [][]int {
	zeroes := []int{}
	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if mat[row][col] == 0 {
				zeroes = append(zeroes, row, col)
			}
		}
	}

	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if mat[row][col] == 0 {
				continue
			}
			lowestDistance := math.MaxFloat64
			for i := 0; i < len(zeroes); i += 2 {
				yDistance := math.Abs(float64(row - zeroes[i]))
				xDistance := math.Abs(float64(col - zeroes[i+1]))
				distance := yDistance + xDistance
				if distance < lowestDistance {
					lowestDistance = distance
					if distance == 1 {
						break
					}
				}
			}
			mat[row][col] = int(lowestDistance)
		}
	}
	return mat
}
