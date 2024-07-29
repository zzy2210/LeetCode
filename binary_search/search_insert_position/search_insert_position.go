package searchinsertposition

/*
35.搜索插入位置 简单

一个是查找，一个是输出位置

先二分，二到最后没用的话就返回 左右
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n-1
	ans := n
	for i < j {
		mid := i + j/2
		if target >= nums[mid] {
			ans = mid
			i = mid + 1
		} else {
			j = mid
		}
	}
	return ans
}
