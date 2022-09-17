package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"unsafe"
)

func init() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	log.Println("init func complete")
}
func main() {
	//t1()
	//t2()
	//t3()
	//t4(1.2)
	//typeOfData()
	processControl()
}

/*变量 声明和变量批量声明*/
func variable() {
	//标准声明
	var name string
	name = "c"
	var age int
	//类型推导  短变量声明
	isOk := true
	//批量声明
	var (
		name1 = 1
		k1    = "d"
	)
	//一次初始化多个变量
	var name2, age2 = "Q1mi", 20
	//匿名变量
	x, _ := foo()
	fmt.Println(x)
	fmt.Println(name2, age2)
	fmt.Println(name, age, isOk)
	fmt.Println(name1, k1)

}
func foo() (int, string) {
	return 10, "Q1mi"
}

/*常量*/
func const_() {
	//iota是go语言的常量计数器，只能在常量的表达式中使用。
	//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	const (
		n1     = iota       //0
		n2                  //1
		_                   //使用_跳过某些值
		n4, n5 = iota, iota //3
		j, k
	)
	const (
		a, b = iota + 1, iota + 2 //1,2  0+1  0+2
		c, d                      //2,3   1+1  1+2
		e, f                      //3,4
	)
	const (
		_  = iota             //1
		KB = 1 << (10 * iota) //1左移十位
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(n1, n2, n4, n5)
	fmt.Println(j, k)
	log.Println(a, b, c, d, e, f)
	log.Println(KB, MB, GB)
	//位运算
	log.Println(1 << 10)
	log.Println(2 >> 3)
	log.Println(1 << (10 * 2))
	log.Println(-1 << 2)
	log.Println(-1 << 10)
	a1, b1 := 1, 2
	var (
		c1 int
		d1 int
	)
	//自反性
	c1 = a1 ^ b1 ^ a1
	d1 = a1 ^ b1 ^ b1
	log.Println(c1, d1)
}

/*运算符*/
func operator() {
	//++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。  golang没有短路与 短路或
	i := 0
	i++
	fmt.Println(i)

}

/*类型断言*/
func typeAssert(i interface{}) {
	if _, ok := i.(int); ok { // 检查类型断言
		return
	}
	switch i.(type) {
	case string:
		fmt.Println("is string")
	case int:
		fmt.Println("is int")
	}
}

/*数据类型*/
func typeOfData() {

	//基本数据类型===========整型
	//整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
	//其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

	//uint	32位操作系统上就是uint32，64位操作系统上就是uint64
	//int	32位操作系统上就是int32，64位操作系统上就是int64
	//uintptr	无符号整型，用于存放一个指针\
	a := 1000
	b := 1
	fmt.Println(a &^ b)
	fmt.Println(a & b) //二进制位与

	c := int64(1)
	d := int32(1)
	e := uintptr(a)
	fmt.Println("ptr", unsafe.Sizeof(e))
	//获取变量c所占的字节
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	//Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字
	//v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。 v := 0o377，代表八进制的 377，相当于十进制的 255。
	//v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。
	//而且还允许我们用 _ 来分隔数字，比如说： v := 123_456 表示 v 的值等于 123456。
	var ba = 010
	ba = 0o11
	var shi = 10
	var shiliu = 0xf
	fmt.Println("进制8/10/16 ", ba, shi, shiliu)
	var v = 123_456_7_89
	var v1 = 0x2222p-2
	//反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	logV := `log_line1
          log line2 inno`
	fmt.Println("v v1", v, v1)
	fmt.Println(logV)
	//字符串的常用操作 join
	s := []string{"1s", "2b", "3c"}
	join := strings.Join(s, " | ")
	//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	//rune类型，代表一个 UTF-8字符。
	//当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
	//Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
	//字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的,字符串是由byte字节组成，
	//所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	//ac:='a'
	//var b rune='d'

	//Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
	fmt.Println("s:", join)
}

/*流程控制*/
func processControl() {
	//if条件判断特殊写法 可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，举个例子：
	/*
		Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。
		同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。
	*/
	if score := 99; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	/*for循环的初始语句和结束语句都可以省略，例如：*/

	ii := 1
	for ; ii < 3; ii++ {
	}
	for ii < 3 {
		ii++
	}
	//无限循环
	for {
		fmt.Println("无限循环")
		break
	}
	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch case_ := 3; case_ {
	case 1, 21, 31:
		fmt.Println("大拇指")
		fallthrough
	case case_:
		fmt.Println("食指")
		//fallthrough
	default:
		fmt.Println("无效的输入！")
	}
	// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	case2 := 12
	switch {
	case false:
		fmt.Println(">20")
	case case2 < 20:
		fmt.Println("<20")
	}

	//for range(键值循环)
	//for i, i2 := range collection {
	//}

	/*
		goto(跳转到指定标签)
		goto语句通过标签进行代码间的无条件跳转。
		goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
		Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：
		//for循环可以通过break、goto、return、panic语句强制退出循环。
	*/
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("j=", j, " ")
			if j == 3 {
				goto tag
			} else {
				goto tag2
			}
		}
	}
tag:
	{
		fmt.Println("goto...tag")
	}
tag2:
	{
		fmt.Println("goto...tag2")
	}
	//
}

/*Go语言fmt.Printf使用指南
fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。
*/
func fmtPrintf() {
	//Print函数直接输出内容，Printf函数支持格式化输
	//Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	//Sprint系列函数会把传入的数据生成并返回一个字符串。
	//Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

	//fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。
	//Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	//bufio.NewReader 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio
	// Fscan系列 fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
	//Sscan系列 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
}
