package main

import (
	"fmt"
	"reflect"
	"strings"
)

//定义全局变量num
var num int64 = 10

func main() {
	a, b := name(1, "a", "b", "c")
	fmt.Println(a, b)
	testGlobalVar()
	//1.全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
	//2/如果局部变量和全局变量重名，优先访问局部变量。
	fmt.Println(calc(12, 22, add))
	//匿名函数
	func(a int, b int) {
		fmt.Println("匿名函数", a, b)
	}(1, 2)
	//闭包 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
	r := A()
	r() //相当于执行匿名函数

	r1 := makeSuffixFunc(".mp4")
	//函数作为变量存在于内存
	s := r1("学习视频")
	fmt.Println(s, r1("ccc"))
	b1 := test("a")
	fmt.Println(b1("b"), b1("C"))
	pf := panicFunc()
	fmt.Println(pf)

}

//1.可变参数可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。 注意：可变参数通常要作为函数的最后一个参数。
//2.Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
/*3.当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
 */
func name(param int, s ...string) (int, []string) {
	fmt.Println(param)
	fmt.Println(s, reflect.TypeOf(s).Kind())
	return 1, nil
}
func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}

type myType func(int, int) int

//定义函数类型
func add(x int, y int) int {
	return x + y
}

//函数作为参数 函数作为返回值
func calc(x, y int, op myType) int {
	return op(x, y)
}
func A() func() {
	name := "little boy"
	return func() {
		fmt.Println("闭包", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//闭包

func test(name string) func(string2 string) string {
	return func(k string) string {
		return name + "ss" + k
	}
}
func panicFunc() (i int) {
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	//panic("函数出错")
	return i
}
