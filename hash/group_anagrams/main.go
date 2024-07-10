package main

import "sort"

// 49. 字母异位词分组
// map 的 k-v k是 某个单词排序后的结果，v放它本来的内容
// 排完之后 直接把每个 v 给 append 进去就完事

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		// 最大的难点居然是这个sort的使用，我真的很少用
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		dict[key] = append(dict[key], str)
	}
	var result [][]string
	for _, array := range dict {
		result = append(result, array)
	}
	return result
}
