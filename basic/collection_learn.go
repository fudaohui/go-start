package main

import "fmt"

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
