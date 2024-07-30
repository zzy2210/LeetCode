package arrowstoburst

import "sort"

/*
 452 用最少数量的箭引爆气球


 题目看上去意思就是，用最少的点覆盖最多的区间
 比如 [1,3],[2,5] 用 2 或则 3 就可以

排序后，一个区域一个区域的找。每个区域用一支箭

*/

func findMinArrowShots(points [][]int) int {
	// 按右端排序
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	count := 0
	i := 0
	for i < len(points) {
		right := points[i][1]
		i++
		// 遍历区间，区间内的每一个左端都小于起始右端
		// 所以必定是重叠
		for i < len(points) && points[i][0] <= right {
			i++
		}
		count++
	}

	return count
}
