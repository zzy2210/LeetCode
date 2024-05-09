package main

import "fmt"

func main() {
	test := Constructor(1)
	test.Put(2, 1)
	fmt.Println(test.Get(2))
	test.Put(3, 2)
	fmt.Println(test.Get(2))
	fmt.Println(test.Get(3))
}

type LRUCache struct {
	count int
	max   int
	head  *Node
	tail  *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

// 初始化
func Constructor(capacity int) LRUCache {
	head := &Node{
		key:   -1,
		value: -1,
	}
	return LRUCache{
		head:  head,
		count: 0,
		max:   capacity,
		tail:  head.next,
	}
}

// 获取
func (this *LRUCache) Get(key int) int {
	node := this.head
	for node != nil {
		if node.key == key {
			// node 脱离原链
			if node.next != nil {
				node.next.prev = node.prev
			}
			node.prev.next = node.next
			if this.tail == node {
				this.tail = node.prev
			}
			// node 头插
			node.next = this.head.next
			node.prev = this.head
			// 如果只有node 一个节点，node脱离后，head.next 是nil
			if this.head.next != nil {
				this.head.next.prev = node
			}
			this.head.next = node
			// 处理尾部
			if node.next == nil {
				this.tail = node
			}

			return node.value
		}
		node = node.next
	}
	return -1
}

// 插入
func (this *LRUCache) Put(key int, value int) {
	node := &Node{
		key:   key,
		value: value,
		prev:  this.head,
		next:  this.head.next,
	}
	if this.head.next != nil {
		this.head.next.prev = node
	}
	this.head.next = node
	if this.count < this.max {
		// 插入
		this.count++
		if this.tail == nil {
			this.tail = node
		}
	} else {
		// 删除
		this.tail = this.tail.prev
		this.tail.next = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
