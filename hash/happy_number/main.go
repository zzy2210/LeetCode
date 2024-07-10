package main

// 202. 快乐数
// 判断非快乐数的方法很简单，如果 某次平方相加的结果出现过，那就一定不是了
func isHappy(n int) bool {
	hash := map[int]bool{}

	for n != 1 && !hash[n] {
		hash[n] = true
		n = step(n)
	}
	return n == 1
}

func step(i int) int {
	sum := 0
	for i > 0 {
		tmp := i % 10
		sum = sum + tmp*tmp
		i = i / 10
	}
	return sum
}
