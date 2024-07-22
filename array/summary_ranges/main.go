package summaryranges

import "fmt"

// 228. 汇总区间
// 题目上来就是已经排序了，并且互不相同
// 直接遍历就好，如果 i +1 ！= a[i+1] 那就说明这一节到它结束
func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var result []string
	tmpres := nums[0]
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			if tmpres == nums[i] {
				result = append(result, fmt.Sprintf("%d", tmpres))

			} else {
				result = append(result, fmt.Sprintf("%d->%d", tmpres, nums[i]))
			}
			break
		}
		if nums[i]+1 != nums[i+1] {
			if tmpres == nums[i] {
				result = append(result, fmt.Sprintf("%d", tmpres))

			} else {
				result = append(result, fmt.Sprintf("%d->%d", tmpres, nums[i]))
			}
			tmpres = nums[i+1]
		}
	}

	return result
}
