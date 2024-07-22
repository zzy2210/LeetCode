package mergeintervals

import "sort"

// 56.合并区间
// 先排序，然后合并就ok
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		// 完全不重合，将pre压进去
		if pre[1] < cur[0] {
			res = append(res, pre)
			pre = cur
		} else {
			//pre = []int{min(pre[0], pre[0]), max(cur[1], pre[1])}
			// 因为是依照[0]排序的，所以pre[0]一定是小于cur[0]的
			pre[1] = max(cur[1], pre[1])
		}
	}
	res = append(res, pre)
	return res
}
