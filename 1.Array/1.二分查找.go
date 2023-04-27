package Array

func search(nums []int, target int) int {
	head := 0
	foot := len(nums)
	var mid int
	for head < foot {
		mid = (head + foot) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			foot = mid
		}
	}
	return -1
}
