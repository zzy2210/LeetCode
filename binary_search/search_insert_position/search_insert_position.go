package searchinsertposition

/*
35.搜索插入位置 简单

一个是查找，一个是输出位置

先二分，二到最后没用的话就返回 左右
*/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r + 1
}
