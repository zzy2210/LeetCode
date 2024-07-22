package twosum

// 1. 两数之和
// 传统 n^2 的解法
func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hash表解法
func twoSumByHash(nums []int, target int) []int {
	// 存放数字
	// k-v : 数字-索引
	hash := make(map[int]int)
	for i, v := range nums {
		// 直接去看表里有没有自己对应的数字
		if p, ok := hash[target-v]; ok {
			return []int{p, i}
		}
		hash[v] = i
	}
	return nil

}
