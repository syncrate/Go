package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func main() {

	p1 := newPerson("small wz", "city1", 25)
	fmt.Printf("newPerson构造函数== %#v \n", p1)
	rc := res{
		id:   20,
		name: "king",
	}
	rc.SetNAme("张三")
	mt := rc.receiverMethod()
	fmt.Println(mt)
	/*需要修改接收者中的值
	接收者是拷贝代价比较大的大对象
	保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。*/

	struct1()
	struct2()
	struct3()
	struct4()
	test1()

}

//Person 结构体Person类型
type Personal struct {
	string
	int
}

type res struct {
	id   int
	name string
}

func (p *res) SetNAme(name string) {
	p.name = name
}

//Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
func (receiver *res) receiverMethod() string {
	id := receiver.id
	name := receiver.name
	return strconv.Itoa(id) + "===" + name
}

/*Go语言中可以使用type关键字来定义自定义类型。*/
type newType int

/*类型别名是Go1.9版本添加的新功能。
类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型*/
type newType1 = newType

func struct1() {
	var a newType
	var b newType1

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}

type Person struct {
	name string
	age  int
	city string
}

func struct2() {
	////只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

	//p1 := Person{}
	//创建指针类型结构体 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
	//p1:=new(Person)
	//取结构体的地址实例化  使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p1 := &Person{}
	p1.city = "cd"
	p1.name = "defaultUser"
	p1.age = 17
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	//匿名结构体
	var user struct {
		Name string
		age  int
	}
	user.Name = "uname"
	user.age = 17
	fmt.Printf("%#v \n", user)
	//没有初始化的结构体，其成员变量都是对应其类型的零值。
	/*
		1.使用键值对初始化
		2.使用值的列表初始化
	*/
	p2 := Person{
		name: "ka",
		age:  12,
		city: "chengdu",
	}
	//必须初始化结构体的所有字段。
	//初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//该方式不能和键值初始化方式混用。
	p3 := &Person{"mi", 16, "h"}
	fmt.Printf("%v  ==", p2)
	fmt.Printf("%+v ==", p2)
	fmt.Printf("%#v ==\n", p2)
	fmt.Printf("%#v \n", p3)

	//空结构体是不占用空间的。
	var v struct{}
	fmt.Println("the size of struct v is :", unsafe.Sizeof(v), "bit") // 0

}

//构造函数
/*Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。*/

func newPerson(name, city string, age_ int) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age_,
	}
}

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	add    Address
	eml    Email
}

func struct3() {

	//结构体的匿名字段
	pp := Personal{
		"小wz",
		18,
	}
	fmt.Println(pp)
	//嵌套结构体

	usr := User{
		Name:   "小王子",
		Gender: "男",
		add: Address{
			"山东",
			"威海",
			"2021",
		},
		eml: Email{"@163.com", "2022"},
	}
	fmt.Printf("%#v\n", usr)
	fmt.Printf("hello world")
	time.Now().Format("2006")
	time.Sleep(time.Second * 1)
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.ant.name)
}

//Dog 狗
type Dog struct {
	Feet int8
	ant  *Animal //通过嵌套匿名结构体实现继承
}

/*
	结构体的继承
*/
func struct4() {
	//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
	d1 := &Dog{
		Feet: 4,
		ant: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang()
	d1.ant.move()

	//JSON序列化：结构体-->JSON格式的字符串

}

//结构体tag

//Student 学生
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"G"` //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func test1() {
	s1 := Student{
		ID:     1,
		Gender: "Male",
		name:   "Shlz",
	}
	data, _ := json.Marshal(s1)
	fmt.Printf("%s  \n", data)
}
