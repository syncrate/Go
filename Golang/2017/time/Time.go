package main

import (
	"fmt"
	"time"
)

func main() {
	timezoneDemo()
	unixTimeDemo()
	timeDemo()
	tickDemo()
	parseDemo()
}
func timeDemo1()  {
	now := time.Now()
	year := now.Year()
	fmt.Println(year,"年")
}
//时区
func timezoneDemo()  {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int(8 * time.Hour.Seconds())

	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing:= time.FixedZone("Beijing Time", secondsEastOfUTC)
	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	fmt.Println("UTC:",secondsEastOfUTC,"北京:",beijing.String(),"本地时间：",newYork)
	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)
	fmt.Println("本地时间：",timeInLocal,"\n",timeInUTC,timeInUTC.Equal(sameTimeInBeijing),timeInUTC.Equal(sameTimeInNewYork))
	
}

//Unix Time是自1970年1月1日 00:00:00 UTC 至当前时间经过的总秒数。下面的代码片段演示了如何基于时间对象获取到Unix 时间。
func unixTimeDemo()  {
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)
}

func timestamp2Time()  {
	// 获取北京时间所在的东八区时区对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 北京时间 2022-02-22 22:22:22.000000022 +0800 CST
	t := time.Date(2022, 02, 22, 22, 22, 22, 22, beijing)

	var (
		sec  = t.Unix()
		msec = t.UnixMilli()
		usec = t.UnixMicro()
	)

	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
	timeObj := time.Unix(sec, 22)
	fmt.Println(timeObj)           // 2022-02-22 22:22:22.000000022 +0800 CST
	timeObj = time.UnixMilli(msec) // 毫秒级时间戳转为时间对象
	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
	timeObj = time.UnixMicro(usec) // 微秒级时间戳转为时间对象
	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
}
func timeDemo()  {
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("现在时间：",format)
	fmt.Println("现在时间：",time.Now())
	/// 当前时间加1小时后的时间
	fmt.Println("一小时之后：",time.Now().Add(1*time.Hour).Format("2006-01-02 15:04:05"))
	//求两个时间之间的差值：
	fmt.Println("一小时之后：",time.Now().Sub(time.Now().Add(1*time.Minute).Add(time.Second).Add(time.Hour))  )
	//判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
	fmt.Println("比较时间：",time.Now()==time.Now())
	fmt.Println("比较时间：",time.Now().Equal(time.Now()))
	fmt.Println("After:",time.Now().After(time.Now().Add(time.Hour))  )
	//如果t代表的时间点在u之后，返回真；否则返回假。
	fmt.Println("Before:",time.Now().Before(time.Now().Add(time.Hour))  )


	
}
//定时器
//使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。
func tickDemo()  {
	//tick := time.Tick(5 * time.Second)
	//for t := range tick {
	//	fmt.Println(t)
	//
	//}
	//如果想格式化为12小时格式，需在格式化布局中添加PM。
	//小数部分想保留指定位数就写0，如果想省略末尾可能的0就写 9。
	now:=time.Now()
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96
	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}


func parseDemo() {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	str := "2022 11 11 11 11 11"
	now:=time.Now()
	parse, err := time.Parse("2006 01 02 15 03 04", str)

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	if err != nil {
		panic(err)
	} else {
		fmt.Println(parse)
	}
	str2:="2022-11-05T11:25:20+08:00"
	t, err := time.Parse(time.RFC3339, str2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(t)
	}
	location, err := time.LoadLocation("Etc/Greenwich")
	if err != nil {
		panic(err)
	}
	inLocation, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/03/15 17:33:20", location)
	if err != nil {
		panic(err)
	}else {
		fmt.Println(inLocation)
		fmt.Println(inLocation.Sub(now.UTC()))
	}
}
















