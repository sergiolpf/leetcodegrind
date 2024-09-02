package main

import "fmt"

func main() {

	// fmt.Printf("%#v", rotate([][]int{
	// 	{5, 1, 9, 11},
	// 	{2, 4, 8, 10},
	// 	{13, 3, 6, 7},
	// 	{15, 14, 12, 16},
	// }))

	matrixDoida := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	traverseReflect(matrixDoida)

	fmt.Println(matrixDoida)

}

/*
[[1,2,3],[4,5,6],[7,8,9]]
[[7,4,1],[8,5,2],[9,6,3]]

Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
*/
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	limI, limJ := (n+1)/2, n/2
	for i := 0; i < limI; i++ {
		for j := 0; j < limJ; j++ {
			temp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}

	return matrix
}

func traverseReflect(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	//reverse

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-j-1]
			matrix[i][len(matrix)-j-1] = temp

		}
	}

}
