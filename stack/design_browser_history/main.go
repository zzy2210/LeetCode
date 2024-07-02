package main

// 一个堆栈
// 下一页就是入栈，上一页就是出栈
// 返回 出栈的内容 同时放入另一个栈内

type BrowserHistory struct {
	preLst []string
	curLst []string
}

func Constructor(homepage string) BrowserHistory {

}

func (this *BrowserHistory) Visit(url string) {

}

func (this *BrowserHistory) Back(steps int) string {

}

func (this *BrowserHistory) Forward(steps int) string {

}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
