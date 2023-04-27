package Array

import "fmt"

func main() {
	//1.二分查找
	//array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//fmt.Println(search(array, 7))
	//2.移除元素
	//nums, val := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	//fmt.Println(removeElement2(nums, val))
	//for _, num := range nums {
	//	fmt.Printf(strconv.Itoa(num))
	//}
	//3、有序数组的平方
	//nums := []int{-5, -3, -2, -1}
	//for _, num := range sortedSquares(nums) {
	//	fmt.Println(num)
	//}
	//4、长度最小的子数组
	//nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	//fmt.Println(minSubArrayLen(11, nums))
	//5、螺旋矩阵
	for _, ints := range generateMatrix(4) {
		for _, num := range ints {
			fmt.Println(num)
		}
	}

}
