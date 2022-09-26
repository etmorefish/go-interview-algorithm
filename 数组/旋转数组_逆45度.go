package main

import "fmt"

/*
将n*n的二位矩阵逆时针旋转45度打印
*/
func rotateArr(arr [][]int) {
	row, col := 0, 0
	l := len(arr)
	// 打印右上部分
	for i := l - 1; i > 0; i-- {
		row = 0
		col = i
		for l > col {
			fmt.Print(arr[row][col], " ")
			row++
			col++
		}
		fmt.Println()
	}
	// 打印左下部分
	for i := 0; i < l; i++ {
		row = i
		col = 0
		for l > row {
			fmt.Print(arr[row][col], " ")
			row++
			col++
		}
		fmt.Println()
	}
}
func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateArr(arr)
}
