package main

import (
	"fmt"
)

func main() {

	//var slice1 []int = make([]int, 5)
	//slice1[0] = 0
	//slice1[1] = 1
	//slice1[2] = 2
	//slice1[3] = 3
	//
	//for i, _ := range slice1 {
	//	slice1[i] = 10
	//}
	//
	//for i, v := range slice1 {
	//	fmt.Printf("index:%d,v:%d\n", i, v)
	//}
	//
	//fmt.Println(len(slice1))
	//fmt.Println(cap(slice1))

	//var arr = [5]int{0, 1, 2, 3, 4}
	//fmt.Printf("%p\n", &arr)
	//sum(arr)

	//-----------test-3
	//var arra [3]int
	//
	//arr1(arra)
	//arrP(&arra)
	//-----------test-4
	//arr[7] = 10
	//fmt.Println(arr)
	//fmt.Println(arr2)

	//同一行多个变量赋值
	//var a int = 10
	//var b int = 11
	//a, b = 20, 31
	//fmt.Printf("%d，%d", a, b)

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	//cap的长度就是开始索引到相关数组结束的长度
	fmt.Printf("%d,%d\n", len(a), cap(a))
}

// 全局变量 数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度
// 全局变量 等同于var arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
var arr2 [5]int

func sum(a [5]int) int {
	fmt.Printf("%p", &a)
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func arr1(a [3]int)  { fmt.Println(a) }
func arrP(a *[3]int) { fmt.Println(a) }

func sliceForRange() {
	str := []string{"a", "b", "c"}
	for _, s1 := range str {
		//for-range结构直接通过value值不能改变数组的值
		s1 = "a1"
		fmt.Printf("s1:%s ", s1)
	}
	fmt.Println(str)
	for ix, _ := range str {
		//for-range结构直接通过数组索引可以改变数组的值
		str[ix] = "cc"
	}
	fmt.Println(str)
}
