package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := book{title: "NoTitle"}
	reflectType(c)
	reflectValue(23.0)
	reflectValue(23)
	vc := reflect.ValueOf(10)
	fmt.Printf("type vc :%T\n", vc) // type c :reflect.Value
	val := 10
	reflectSetValue1(val)
	reflectSetValue2(val)
}

//1.在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的
//2.reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
type book struct{ title string }

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v value:%v name:%v kind:%v  \n", v, reflect.ValueOf(x), v.Name(), v.Kind())
}

//1.在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
//2.因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型.
//3.但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）

//通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int16:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
//而反射中使用专有的Elem()方法来获取指针对应的值。
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
