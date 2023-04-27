package Array

/*
*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [3,-1,0,16,100] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

要求时间复杂度On
*/
func sortedSquares(nums []int) []int {
	low := 0
	high := len(nums) - 1
	var newArray = make([]int, len(nums))
	index := high
	for low <= high {
		if nums[low]*nums[low] <= nums[high]*nums[high] {
			newArray[index] = nums[high] * nums[high]
			high--
		} else {
			newArray[index] = nums[low] * nums[low]
			low++
		}
		index--
	}
	return newArray
}
