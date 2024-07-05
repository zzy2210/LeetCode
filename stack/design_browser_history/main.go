package main

import "fmt"

// 一个堆栈
// 下一页就是入栈，上一页就是出栈
// 返回 出栈的内容 同时放入另一个栈内

func main() {
	bro := Constructor("leetcode.com")
	bro.Visit("google.com")
	bro.Visit("facebook.com")
	bro.Visit("youtube.com")
	fmt.Println(bro.Back(1))
	fmt.Println(bro.Back(1))
	fmt.Println(bro.Back(1))
	bro.Visit("linkedin.com")
	fmt.Println(bro.Forward(2))
	fmt.Println(bro.Back(2))
	fmt.Println(bro.Back(7))

}

type BrowserHistory struct {
	preLst []string // 前进队列
	curLst []string // 后退队列
	url    string   // 当前
}

// 初始化浏览器
func Constructor(homepage string) BrowserHistory {

	return BrowserHistory{
		preLst: []string{},
		curLst: []string{},
		url:    homepage,
	}
}

// 跳转页面
// 会清空前面的历史记录
func (this *BrowserHistory) Visit(url string) {
	this.curLst = append(this.curLst, this.url)
	this.preLst = []string{}
	this.url = url
}

func (this *BrowserHistory) Back(steps int) string {
	if len(this.curLst) == 0 {
		return this.url
	}
	if steps >= len(this.curLst) {
		steps = len(this.curLst) // step = 0
	}

	tmpList := this.curLst[len(this.curLst)-steps:]    //  [1-0=1:]
	this.curLst = this.curLst[:len(this.curLst)-steps] // [:1-0=0]
	for i := range tmpList {
		this.preLst = append(this.preLst, this.url)
		this.url = tmpList[i]
	}
	return this.url
}

func (this *BrowserHistory) Forward(steps int) string {
	if len(this.preLst) == 0 {
		return this.url
	}
	if steps >= len(this.preLst) {
		steps = len(this.preLst)
	}
	tmpList := this.preLst[len(this.preLst)-steps:]
	this.preLst = this.preLst[:len(this.preLst)-steps]
	for i := range tmpList {
		this.curLst = append(this.curLst, this.url)
		this.url = tmpList[i]
	}
	return this.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
