package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	/*Go语言中 `map`的定义语法如下：  map[KeyType]ValueType*/
	//map也支持在声明的时候填充元素，例如：
	map_ := map[string]int{"1": 1, "license": 12315}
	//map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	fmt.Println("map:", map_)
	map1 := make(map[string]string, 3)
	//其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量
	fmt.Println("map1分配内存:", map1)
	//how to use
	m := make(map[string]int, 8)
	m["张三"] = 90
	m["李四"] = 100
	fmt.Println(m)
	delete(m, "张三") //如果key不存在则 不删除
	fmt.Println("删除元素后的M:", m)
	//遍历map时的元素顺序与添加键值对的顺序无关。 无序
	//按照指定顺序遍历map？

	//值为切片类型的map
	sliceMap := make(map[string][]string, 3)
	fmt.Println("sliceMap:", sliceMap)
	key := "cn"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	sliceMap["en"] = make([]string, 5, 5)
	copy(sliceMap["en"], value)
	fmt.Println(sliceMap, ok)
	sliceMap["en"][0] = "Beijing"
	//Go语言中使用for range遍历map。 k,v
	for s, strings := range sliceMap {
		fmt.Println("start==", s, strings, "==end")
	}
	//使用delete()内建函数从map中删除一组键值对，
	delete(sliceMap, "en")
	fmt.Println(sliceMap)
	//按照指定顺序遍历map
	sortMap()
	notBuiltinFunction()
}
func notBuiltinFunction() {
	fmt.Println("不是一个内置方法")
}
func sortMap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Print(key, " ", scoreMap[key], "  ")
	}
	fmt.Println()
}
