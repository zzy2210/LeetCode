package evaluatereverse

import "strconv"

/*
	150 逆波兰表达式求值

	感觉还是堆栈哈

	遇到表达式就给栈里的两个值拿出来，算好了再压进去


*/

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			switch token {
			case "+":
				stack = append(stack[:len(stack)-2], v1+v2)
			case "-":
				stack = append(stack[:len(stack)-2], v1-v2)
			case "*":
				stack = append(stack[:len(stack)-2], v1*v2)
			case "/":
				stack = append(stack[:len(stack)-2], v1/v2)
			default:
				// err
			}
		} else {
			stack = append(stack, num)
		}
	}

	return stack[0]
}
