package main

import (
	"strings"
)

// 290. 单词规律

/*
某座映射规律，那也做map1
这题与 同构字符串 其实没有区别
*/

func wordPattern(pattern string, s string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	// 拆分 s
	s1 := strings.Split(s, " ")
	if len(pattern) != len(s1) {
		return false
	}
	for i := range pattern {
		x, y := pattern[i], s1[i]
		if _, ok := p2s[x]; !ok {
			p2s[x] = y
		}
		if _, ok := s2p[y]; !ok {
			s2p[y] = x
		}
		if p2s[x] != y || s2p[y] != x {
			return false
		}
	}
	return true
}
