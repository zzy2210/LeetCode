package insertinterval

/*

57 插入区间  中等

一个已经排序了的区间,然后插入新区间.
对新的区间做一遍合并,然后返回
感觉其实就是合并区间改一改

1,3  2,5
重叠的方式是right > newLeft

看了一个想法 是三段

第一段:  i[1] < add[0] 完全在左边的部分,无脑加入
第二段:  i[0] <= add[1] , 到了这一步已经暗含了  i[1] < add[0]  所以是完全在中间的部分,做合并行为,并且加入
第三段: i[0]> a[0] 完全在右边的部分,无脑加入

*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	n := len(intervals)
	i := 0

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)
	for i < n {
		result = append(result, intervals[i])
		i++
	}
	return result
}
