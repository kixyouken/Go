package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("当前时间带毫秒：", t)

	fmt.Println("当前时间戳：", t.Local().Unix())
	// PHP 格式化时间是 Y-m-d H:i:s
	// GO 格式化时间是 2006-01-02 15:04:05
	fmt.Println("格式化当前时间戳：", t.Local().Format("2006-01-02 15:04:05"))

	date := "2023-01-01 10:00:00"
	format_time, _ := time.Parse("2006-01-02 15:04:05", date)
	fmt.Println("日期转时间戳后：", format_time.Unix())

	fmt.Println("本地当前年份方法1：", t.Local().Year())
	fmt.Println("本地当前年份方法2：", t.Local().Format("2006"))

	fmt.Println("本地当前月份英文1：", t.Local().Month())
	fmt.Println("本地当前月份英文2：", t.Local().Month().String())
	fmt.Println("本地当前月份数字：", t.Local().Format("01"))

	fmt.Println("本地当前日期方法1：", t.Local().Day())
	fmt.Println("本地当前日期方法2：", t.Local().Format("02"))

	fmt.Println("本地当前时间时1：", t.Local().Hour())
	fmt.Println("本地当前时间时2：", t.Local().Format("15"))

	fmt.Println("本地当前时间分1：", t.Local().Minute())
	fmt.Println("本地当前时间分2：", t.Local().Format("04"))

	fmt.Println("本地当前时间秒1：", t.Local().Second())
	fmt.Println("本地当前时间秒2：", t.Local().Format("05"))

	fmt.Println("本地当前时间纳秒：", t.Local().Nanosecond())

	// 同时输出年月日
	fmt.Println(t.Local().Date())

	fmt.Printf("今天是当前年份的第 %d 天\n", t.Local().YearDay())

	fmt.Printf("今天是当前月份的第 %d 天\n", t.Local().Day())

	fmt.Println("今天是星期几：", time.Weekday(time.Now().Weekday()))
}
