package isomorphicstrings

//205. 同构字符串

/*
不管是某种映射
s到 t 与 t 到 s
必定每个字符都有唯一的映射对象
那就建立 两个map，便利地往里面写
k-v 是 值与他的映射

如果出现了 已经有映射并且与新映射值不同的情况，就说明不对
返回错误

*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if _, ok := s2t[x]; !ok {
			s2t[x] = y
		}
		if _, ok := t2s[y]; !ok {
			t2s[y] = x
		}
		if s2t[x] != y || t2s[y] != x {
			return false
		}
	}
	return true
}
