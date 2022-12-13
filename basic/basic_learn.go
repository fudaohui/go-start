package basic

import (
	"fmt"
	"io"
	"log"
)

func pointerTest1() {
	var i1 = 5
	fmt.Printf("location-0:%p\n", &i1)
	var intP *int = &i1
	fmt.Printf("location-1:%p\n", intP)
	fmt.Println(*intP == i1)
}
func poinerTest2() {

	//s := "good bye"
	//var strP *string = &s
	//
	////将指针strP 指向地址值修改为ss
	//*strP = "ss"
	//
	//fmt.Print(s)

	var p1 *int = nil
	*p1 = 0
}

func controllStruture1(x int) int {
	//fmt.Println(runtime.GOOS)

	if x < 0 {
		return -x
	}
	return x
}

func controlStruture2(a int) {

	//初始化和判断在同一行
	if v := controllStruture1(a); v > 100 {
		fmt.Println(v + 10)
	}
}

func multiReturnDataTest() {

	//error 模型
	//s := "hello"
	//_, err := strconv.Atoi(s)
	//if err != nil {
	//	return
	//
	//}

	okPatterm(10)

}

func okPatterm(v int) (d1 int, ok bool) {

	if v > 0 {
		return v + 10, true
	}
	return 0, false
}

// 简单的for快捷键是fori
func forTest() {

	//str := "日本語"
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c\n", str[i])
	//}

	//for i := 0; i < 5; i++ {
	//	for j := 0; j <= i; j++ {
	//		fmt.Print("G")
	//	}
	//	fmt.Println()
	//}

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func bitReverse() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%b\n", ^i)
	}
}

func fortest2() {

	for true {
		fmt.Println("*")
	}

}

func forRangeTest(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func breakTest() {

	var a int = 10

	for {
		a -= 1
		if a < 6 {
			break
		}
		fmt.Print(a)
	}

}

// Func1 测试
func Func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
