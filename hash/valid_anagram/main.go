package main

// 242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	d1 := make(map[byte]int)

	for i := range s {
		d1[s[i]]++
	}

	for i := range t {
		if _, ok := d1[t[i]]; !ok {
			return false
		}
		d1[t[i]]--
		if d1[t[i]] == 0 {
			delete(d1, t[i])
		}
	}
	return true
}
