package main

import (
	"flag"
	"fmt"
)

func main() {
	kfc()

}

func kfc ()  {
	var n2 string
	var n3 string
	flag.StringVar(&n2, "n2", "", "疯狂")
	flag.StringVar(&n3, "n3", "", "星期四")
	flag.Parse()
	args := flag.Args()
	if args[0]=="crazy" &&args[1]=="thursday" {
		fmt.Println("Give Me 50 RMB")
	}else {
		fmt.Println("命令错误")
	}

}
