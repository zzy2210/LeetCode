package longestsubstring

func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{}
	n := len(s)
	right, maxLen := 0, 0

	for i := 0; i < n; i++ {
		// 走到这一步的时候，已经做过 一些查找了
		if i != 0 {
			// 已经到了当前窗口最大，左边界右移
			delete(dict, s[i-1])
		}
		for right < n && dict[s[right]] == 0 {
			dict[s[right]]++
			right++
		}
		// 一个完整窗口出来了，算窗口最大值
		maxLen = max(maxLen, right-i)

	}
	return maxLen
}
