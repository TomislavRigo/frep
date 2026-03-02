package matcher

func initlizeMatrix(sourceLen, targetLen int) [][]int {
	matrix := make([][]int, sourceLen+1)
	for i := range matrix {
		matrix[i] = make([]int, targetLen+1)
	}

	for i := range len(matrix[0]) {
		matrix[0][i] = i
	}

	for i := range len(matrix) {
		matrix[i][0] = i
	}

	return matrix
}

func CalculateDistance(source, target string) float64 {
	matrix := initlizeMatrix(len(source), len(target))

	for y := 1; y < len(matrix[0]); y++ {
		for x := 1; x < len(matrix); x++ {
			cost := 0
			if source[x-1] != target[y-1] {
				cost = 1
			}

			del := matrix[x-1][y] + 1
			ins := matrix[x][y-1] + 1
			sub := matrix[x-1][y-1] + cost

			matrix[x][y] = min(del, ins, sub)
		}
	}

	distance := matrix[len(source)][len(target)]
	return (1 - float64(distance)/max(float64(len(source)), float64(len(target))))
}
