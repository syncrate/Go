package main

import (
	d "database/sql"
	"github.com/q1mi/hello"
)

func main() {
	/*1  如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，
	那么标识符必须是对外可见的（public）。在Go语言中是通过标识符的首字母大/小写来控制标识符的对外可见（public）/不可见（private）的。
	在一个包内部只有首字母大写的标识符才是对外可见的。
	*/

	/*2  import name：引入的包名，通常都省略。默认值为引入包的包名。
	path/to/package：引入包的路径名称，必须使用双引号包裹起来。
	*/
	/*3  如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。
	一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。
	被匿名引入的包中的init函数将被执行并且仅执行一遍。*/
	d.Drivers()
	/*
		module holiday
		go 1.16
		require github.com/q1mi/hello v0.1.0 // indirect

		go mod
		module holiday：定义当前项目的导入路径
		go 1.16：标识当前项目使用的 Go 版本

		行尾的indirect表示该依赖包为间接依赖，说明在当前程序中的所有 import 语句中没有发现引入这个包。
		另外在执行go get命令下载一个新的依赖包时一般会额外添加-u参数，强制更新现有依赖。
		//如果事先知道依赖包的具体版本号，可以直接在go.mod中指定需要的版本然后再执行go mod download下载。
	*/
	hello.SayHi("Ni")

}
