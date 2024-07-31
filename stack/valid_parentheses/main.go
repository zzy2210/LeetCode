package validparentheses

/*
	20 有效的括号
	一言堆栈
	来了数据和栈顶的对比，可以匹配就消了，匹配不了就入栈
	跑完后如果栈内为空 就ok
*/

func isValid(s string) bool {
	stack := []byte{}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, byte(v))
			continue
		}
		i := len(stack) - 1
		if (stack[i] == '(' && v == ')') || (stack[i] == '[' && v == ']') || (stack[i] == '{' && v == '}') {
			stack = stack[:i]
		} else {
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}
