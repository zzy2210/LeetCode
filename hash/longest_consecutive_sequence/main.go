package longestconsecutivesequence

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := map[int]bool{}
	for _, v := range nums {
		numSet[v] = true
	}
	result := 0
	for num := range numSet {
		if !numSet[num-1] {
			current := num
			cruuentStreak := 1
			for numSet[current+1] {
				current++
				cruuentStreak++
			}
			if result < cruuentStreak {
				result = cruuentStreak
			}
		}
	}
	return result

}
