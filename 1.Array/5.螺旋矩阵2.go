package main

/*
*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	i, j, num := 0, 0, 1
	for k := 0; k < n/2; k++ {
		for j = k; j < n-k-1; j++ {
			result[i][j] = num
			num++
		}
		for i = k; i < n-k-1; i++ {
			result[i][j] = num
			num++
		}
		for ; j > k; j-- {
			result[i][j] = num
			num++
		}
		for ; i > k; i-- {
			result[i][j] = num
			num++
		}
		i++
	}
	if n%2 == 1 {
		result[n/2][n/2] = n * n
	}
	return result
}
