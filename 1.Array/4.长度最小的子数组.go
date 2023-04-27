package Array

/*
*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
//暴力求法On2
//func minSubArrayLen(target int, nums []int) int {
//	minLen := 0
//	for i, num := range nums {
//		sum := num
//		j := i + 1
//		for j < len(nums) && sum < target {
//			sum += nums[j]
//			j++
//		}
//		if (sum >= target) && (minLen == 0 || j-i < minLen) {
//			minLen = j - i
//		}
//	}
//	return minLen
//}

//func minSubArrayLen(target int, nums []int) int {
//	minLen := 0
//	total := 0
//	for j := 0; j < len(nums); j++ {
//		i := 0
//		total += nums[j]
//		sum := total
//		if sum < target {
//			continue
//		}
//		for i < j {
//			sum -= nums[i]
//			if sum < target {
//				break
//			}
//			i++
//		}
//		if j-i < minLen || minLen == 0 {
//			minLen = j - i + 1
//		}
//	}
//	return minLen
//}

// 双指针（滑动窗口）
func minSubArrayLen(target int, nums []int) int {
	minLen := 0
	sum := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if minLen == 0 || j-i+1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	return minLen
}
