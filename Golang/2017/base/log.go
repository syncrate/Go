package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		const (
		    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
		    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
		    Ldate         = 1 << iota     // 日期：2009/01/23
		    Ltime                         // 时间：01:23:23
		    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
		    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
		    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
		    LUTC                          // 使用UTC时间
		    LstdFlags     = Ldate | Ltime // 标准logger的初始值
		)
	*/

	//func New(out io.Writer, prefix string, flag int) *Logger
	//New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
	a, b, c := 0x12, 0o12, 0b11
	fmt.Println(a, b, c)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[前缀]")
	log.Println("test")
	logFile, err := os.OpenFile("2017/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0064)
	if err != nil {
		fmt.Println("open log file error: ", err)
	}
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
	//log.SetOutput(logFile)
	fmt.Println(logFile)
	logger.Println("this is a normal log")
	v := "very import"
	logger.Printf("this is a %s log ", v)
	//log.Fatalf("出发Fatalf")
	//log.Panicln("触发panic")
	fmt.Println("打印================")
	//配置logger

}
