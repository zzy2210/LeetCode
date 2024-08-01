package basiccalculator

/*

224 基本计算器 困难


*/

func calculate(s string) int {
	result := 0
	// 表示一个个闭合括号， +1 或则 -1 代表结果是正还是负
	opt := []int{1}
	// 只有 +1 ,-1情况的标志位
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			// 跳过
			i++
		// 根据操作符的情况，取正负
		case '+':
			sign = opt[len(opt)-1]
			i++
		case '-':
			sign = -opt[len(opt)-1]
			i++
		case '(':
			opt = append(opt, sign)
			i++
		case ')':
			// 闭合了一个，取出
			opt = opt[:len(opt)-1]
			i++
		default:
			nums := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				nums = nums*10 + int(s[i]-'0')
			}
			result += sign * nums
		}

	}
	return result
}

func calculate_a(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
