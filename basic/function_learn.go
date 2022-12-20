// Package basic go语言中的函数学习.
// 重点是：求解大事需要拆解为小事，小事解决了大事自然而然就解决了。小事求解到一定的程度必须有退出条件
package main

import (
	"fmt"
	"time"
)

/*
Fibonacci 斐波那契数列数列计算，使用递归的方式。斐波那契数列形式如下：
1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
递归的思路就是如求解大事需要拆解为小事，小事解决了大事自然而然就解决了
*/
func Fibonacci(n int) (res int) {
	if n <= 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Print10to1  10到n 的数据打印
func Print10to1(i int) {
	if i > 10 {
		return
	}
	Print10to1(1 + i)
	fmt.Printf("%d ", i)
}

// Function_param 函数参数，回调的使用
func Function_param(n int, f func(n1 int, n2 int)) {
	f(n+1, n+2)
}

func Callback(n1 int, n2 int) {
	fmt.Printf("%d-%d", n1, n2)
}

// Noname 匿名函数的使用
func Noname() {
	//不想起名的函数g实际上起了名字g
	g := func(n int) { fmt.Println(n * 2) }
	//变量参数当函数使用，相当于函数名为g了？？
	g(3)
}

// ExecTimeCacl 函数的执行时间
func ExecTimeCacl() {

	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
