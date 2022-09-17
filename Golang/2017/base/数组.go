package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	log.Println("Start Main func")
	array()

}

func init() {
	log.Println("Init Complete")
}

/*数组*/
func array() {
	//定义一个长度为3的数组
	//var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。
	var a [3]int
	fmt.Println("数组a:", a, reflect.TypeOf(a).Kind())
	//初始化
	var arr1 [3]int
	var arr2 = [3]int{1, 2}
	fmt.Println(arr1, arr2)
	//一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	var arr3 = [...]int{1, 3}
	//我们还可以使用指定索引值的方式来初始化数组，例如
	var arr4 = [...]int{1, 2, 3: 4, 4: 11, 42}
	fmt.Printf("%T ,%T \n", arr3, arr4)
	fmt.Println("arr3 arr4:", arr3, arr4)
	//变量
	for index, value := range arr4 {
		fmt.Println("K:V", index, value)
	}
	for i := 0; i < len(arr4); i++ {
		fmt.Print("arr4:", arr4[i], "  ")
	}
	c := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"}, {"Admin"}}
	//二维数组
	fmt.Println("a:=", c)
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	//**注意：** 多维数组**  只有第一层 可以使用...来让编译器推导数组长度。例如：
	array2 := [...][3]string{
		{"2d", "2x"},
		{"xa", "sad"},
	}
	fmt.Println("array2:", array2)

	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

	//[n]*T表示指针数组，*[n]T表示数组指针 。
	t1 := [3]*string{}
	t2 := new(*[3]string)
	fmt.Println(t1, t2)

	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	ar := [3]int{10, 20, 30}
	modifyArray(&ar) //在modify中修改的是a的副本x
	fmt.Println(ar)  //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]

}
func modifyArray(x *[3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
