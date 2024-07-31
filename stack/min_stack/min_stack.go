package minstack

/*
	155 最小栈
	看上去就是一个栈+最小值搜索

	那直接定义一个常量放最小值不就ok了
	但是这样有个问题就是，如果最小值被pop了，怎么办


	那就不用常量存储，用栈。每次push的时候对比最小值栈的顶部和入栈内容。

	如果最小不变，辅助栈再压入一个最小

	如果最小变了，辅助栈压这个。总之就是记录每一个实际栈的那一点，对应的最小值的栈

	就是有点抽象，栈里放栈

*/

type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: []int{},
		stack:    []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
