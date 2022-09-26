package main

// 顺时针旋转 90 度。
func rotate(matrix [][]int) {
	// 1. 先沿对角线翻折
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 2. 然后翻转每一行
	for _, row := range matrix {
		i, j := 0, len(row)-1
		for i < j {
			row[i], row[j] = row[j], row[i]
			i++
			j--
		}
	}
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}
