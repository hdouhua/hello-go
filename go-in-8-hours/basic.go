package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	println("hello, world!")
	// fmt.Println("test import statement")

	// ----- string

	// 多行字符串
	println(`line 1
    line 2
    line 3`)

	// 编码相关
	println(len("你好"))
	println(len("你好abc"))
	println(utf8.RuneCountInString("你好"))
	println(utf8.RuneCountInString("你好abc"))

	// ----- array

	a1 := [3]int{3, 2, 1} // 区别于 a1 := []int{3, 2, 1} ，后者是切片
	// ==> a1 := [...]int{3, 2, 1} // ... 代表数组长度自动推导
	fmt.Printf("a1: %v, len: %d, cap: %d, 1st emlement: %d\n", a1, len(a1), cap(a1), a1[0])

	var a2 [3]int
	// ==> var a2 = [3]int{0, 0, 0}
	// ==> var a2 = [...]int{0, 0, 0}
	fmt.Printf("a2: %v, len: %d, cap: %d\n", a2, len(a2), cap(a2))

	// ----- slice

	s1 := []int{}
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 5)
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 3)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 3, 4)
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := make([]int, 3)
	// ==> s3 := make([]int, 3, 3)
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	// ----- for

	println("--- for while ----")
	index := 0
	for {
		if index >= len(a1) {
			break
		}
		fmt.Printf("a1[%d]: %d\n", index, a1[index])
		index++
	}
	println("--- for i----")
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%d]: %d\n", i, a1[i])
	}
	println("---- for range ---")
	for i, v := range a1 {
		fmt.Printf("a1[%d]: %d\n", i, v)
	}
	println("a1 contains: ")
	for _, v := range a1 {
		println(v)
	}

	// ----- if
	if isTriangle(3, 5, 4) {
		println("this is a triangle")
	} else {
		println("this is NOT a triangle")
	}

	// ----- switch

}

// ----- if
func isTriangle(a, b, c int) bool {
	// 以下 t1, t2, t3 是局部变量，作用域只在 if - else 上下文中
	if t1 := a + b; t1 > c {
		if t2 := a + c; t2 > b {
			if t3 := b + c; t3 > a {
				fmt.Printf("the sum of any two sides of a triangle: %d, %d, %d\n", t1, t2, t3)
				return true
			}
		}
	}
	return false
}
