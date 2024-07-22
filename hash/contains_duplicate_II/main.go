package containsduplicateii

// 219. 存在重复元素 II

// 先搞个hash，遍历写入
// 如果 hash到相同了，就计算 如果对了返回true
func ContainsNearbyDuplicate(nums []int, k int) bool {
	// 值- 索引
	hash := map[int]int{}

	for i, v := range nums {
		if j, ok := hash[v]; ok && abs(i, j) <= k {
			return true
		}
		hash[v] = i
	}

	return false

}

func abs(i, j int) int {
	if i >= j {
		return i - j
	}
	return j - i
}
