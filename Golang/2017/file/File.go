package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//openFileRead()
	readFile()
	cRead()
}
func openFileRead() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./2017/file/read.go")
	defer file.Close()
	if err != nil {
		panic(err)
	}
}
//循环读取
func readFile()  {
	file, err := os.Open("./2017/file/File.go")
	if err != nil {
		panic(err)

	}
	defer file.Close()
	//它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。 举个例子
	temp := make([]byte, 128)
	n, err := file.Read(temp)
	if err== io.EOF{
		fmt.Println("读完了")
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(temp[:]))

}
func cRead()  {
	file, err := os.Open("./2017/file/File.go")
	if err != nil {
		fmt.Println("Err:",err)
		return
	}
	defer file.Close()

}