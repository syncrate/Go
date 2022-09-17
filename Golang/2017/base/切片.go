package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*声明切片类型的基本语法如下：
	var name []T
	*/
	a := [6]int{1, 2, 3, 4, 5, 6}
	var b []int
	fmt.Println(reflect.TypeOf(b).Kind(), b)
	/*切片拥有自己的长度和容量，我们可以通过使用内置的`len()`函数求长度，使用内置的`cap()`函数求切片的容量。*/
	/*一个索引范围（左包含，右不包含） 左闭右开区间
	a[2:]  // 等同于 a[2:len(a)]
	a[:3]  // 等同于 a[0:3]
	a[:]   // 等同于 a[0:len(a)]
	完整切片表达式 a[low : high : max]

	*/
	s := a[3:4]
	fmt.Println(reflect.TypeOf(a).Kind(), a, len(a), cap(a), " 切片S:", s, len(s), cap(s))
	/*
		make([]T, size, cap)
	*/
	c := make([]int, 2, 10)
	fmt.Println(c, len(c), cap(c))
	//要检查切片是否为空，请始终使用`len(s) == 0`来判断，而不应该使用`s == nil`来判断。
	s1 := make([]int, 3)
	s2 := s1 //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 666
	fmt.Println(s1, s2)
	//切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。
	for index, value := range s2 {
		fmt.Print(index, value, "======")
	}
	fmt.Println()
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], " ")
	}
	fmt.Println()
	//Go语言的内建函数`append()`可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
	//：**通过var声明的零值切片可以在`append()`函数直接使用，无需初始化。
	s2 = append(s2, 3, 4, 5)
	fmt.Println(s2)
	//由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

	//经典的切片扩容问题
	aa := make([]int, 3, 10)
	aa[0], aa[1], aa[2] = 1, 2, 3
	bb := aa[:]
	fmt.Println(len(bb), cap(bb))
	aa = append(aa, 4)
	bb = append(bb, 5)
	aa = append(aa, 6)
	bb = append(bb, 51)
	aa[0] = 1000
	fmt.Println(aa, bb)

	fmt.Println("----分割线-----")

	copyA := []int{1, 2, 3}
	copyB := make([]int, 5, 5)
	copy(copyB, copyA)
	fmt.Println(copyA, copyB)

}
