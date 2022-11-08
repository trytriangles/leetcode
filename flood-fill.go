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

func Test_floodFill(t *testing.T) {
	cases := []testCase{
		{
			Input: Input{
				Image: [][]int{{0, 0, 0}, {0, 0, 0}},
				Sr:    0,
				Sc:    0,
				Color: 0,
			},
			Expected: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
		{
			Input: Input{Image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				Sr:    1,
				Sc:    1,
				Color: 2,
			},
			Expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for index, testCase := range cases {
		result := floodFill(testCase.Input.Image, testCase.Input.Sr, testCase.Input.Sc, testCase.Input.Color)
		if len(result) != len(testCase.Expected) {
			t.Fatal("row count mismatch")

		}
		for i := 0; i < len(result); i++ {
			if len(result[i]) != len(testCase.Expected[i]) {
				t.Fatalf("length mismatch on row %v", i)
			}
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != testCase.Expected[i][j] {
					t.Fatalf("testcase %v: result[%v][%v]=%v, expected %v", index, i, j, result[i][j], testCase.Expected[i][j])
				}
			}
		}
	}
}