// Package basic go语言中的函数学习
package basic

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
