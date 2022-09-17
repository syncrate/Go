package main

import "fmt"

func main() {
	example1()
	newVar()
}

func example1() {
	/*Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）。*
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
	Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
	Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。
	*/
	a := 10
	b := &a
	fmt.Printf("valueA:%d  pointA:%p \n", a, &a)
	fmt.Printf("b:%p type: %T \n", b, b)
	c := *b
	fmt.Printf("value:%v , type:%T \n", c, c)

}
func newVar() {
	/*Type表示类型，new函数只接受一个参数，这个参数是一个类型
	*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	*/
	a := new(*int)
	b := new(int)
	fmt.Println(*a, *b)
	c := make(map[string]string, 10)
	fmt.Println(c)
	//二者都是用来做内存分配的。
	//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
