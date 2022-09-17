package main

import "fmt"

func main() {
	new(Bird).Sing()
	var c Cat
	var d Dogs
	makeHungry(c)
	makeHungry(d)
	d.Move()
	c.Move()
	// 多种类型实现同一接口
	fmt.Println("Interface")
	var s Sayer
	s = Dogs{}
	s.Say()
	s = Cat{}
	s.Say()
	testInterface()

}

/*每个接口类型由任意个方法签名组成，接口的定义格式如下：
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
*/

/*
2.接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，
那么我们就称它实现了这个接口。
*/

//Singer 接口
type Singer interface {
	Sing()
}
type Bird struct {
}

func (b Bird) Sing() {
	fmt.Println("====implement Sing function")
}

type Sayer interface {
	Say()
}
type Dogs struct{}
type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}
func (c Dogs) Say() {
	fmt.Println("汪汪汪~")
}
func makeHungry(s Sayer) {
	s.Say()
}

type Mover interface {
	Move()
}

// Move 值接收者实现接口
func (d Dogs) Move() {
	//panic("implement me")
	fmt.Println("Dog Can Move")
}

// Move 指针接收者实现接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

//// 下面的代码无法通过编译
//var c2 = Cat{} // c2是Cat类型
//x = c2         // 不能将c2当成Mover类型
//由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。
//但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

//一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// Haier 海尔洗衣机
type Haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h Haier) wash() {
	fmt.Println("洗刷刷")
}

//接口与接口之间可以通过互相嵌套形成新的接口类型，例如Go标准库io源码中就有很多接口之间互相组合的示例。
//1.空接口作为函数的参数  使用空接口实现可以接收任意类型的函数参数。
//2.使用空接口实现可以保存任意值的字典。
// 空接口作为map值
//var studentInfo = make(map[string]interface{})
//studentInfo["name"] = "沙河娜扎"
//studentInfo["age"] = 18
//studentInfo["married"] = false
//fmt.Println(studentInfo)

// Runner *由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。也就是说接口值由“类型”和“值”组成，//
//鉴于这两部分会根据存入值的不同而发生变化，
//我们称之为接口的动态类型和动态值。*/
type Runner interface {
	Run()
}
type SmallCar struct {
}
type LargeCar struct {
}

func (s *SmallCar) Run() {
	fmt.Println("小汽车")
}
func (s *LargeCar) Run() {
	fmt.Println("大汽车")
}
func testInterface() {
	var r Runner
	fmt.Println(r == nil)
	//	??注意：我们不能对一个空接口值调用任何方法，否则会产生panic。
	//	r.Run()
	r = &SmallCar{}
	r.Run()
	//此时，接口值r的动态类型会被设置为*SmallCar，动态值为结构体变量的拷贝。

	//而想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。
	//x：表示接口类型的变量
	//T：表示断言x可能是的类型。
	var n Mover = &Dogs{}
	v, ok := n.(*Cat)
	if ok {
		fmt.Println("断言成功", v)
	} else {
		fmt.Println("断言失败")
	}

}
