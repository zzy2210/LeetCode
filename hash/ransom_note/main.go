package main

/*
383.赎罪信

字典法 创建一个map
发现将 ransomNote 的每个字符往里面写
出现一次就是 dic[a] = 1
再一次就再加

然后遍历b，遍历到某个存在的数值就减一，为0 则删除

之后查看字典，如果字典为空，则说明有效
*/
func canConstruct(ransomNote string, magazine string) bool {
	dict := make(map[rune]int)
	for _, v := range ransomNote {
		dict[v]++
	}

	for _, v := range magazine {
		if _, ok := dict[v]; ok {
			dict[v]--
			if dict[v] == 0 {
				delete(dict, v)
			}
		}
	}
	return len(dict) == 0
}
