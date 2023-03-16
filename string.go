package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "     this is string "
	fmt.Println(str)

	// 判断str是否有前缀字符串prefix。
	boolean := strings.HasPrefix(str, "this")
	fmt.Println("是否有前缀：", boolean)

	// 判断str是否有后缀字符串suffix。
	boolean = strings.HasSuffix(str, "this")
	fmt.Println("是否有后缀：", boolean)

	// 判断字符串str是否包含子串substr。
	boolean = strings.Contains(str, "is")
	fmt.Println("是否包含：", boolean)

	// 判断字符串str是否包含字符串chars中的任一字符。
	boolean = strings.ContainsAny(str, "is array")
	fmt.Println("是否包含任一字符：", boolean)

	// 返回字符串str中有几个不重复的sep子串。
	count := strings.Count(str, "s")
	fmt.Println("子串的数量：", count)

	// 子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	count = strings.Index(str, "s")
	fmt.Println("子串第一次出现的位置：", count)

	// 子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	count = strings.LastIndex(str, "s")
	fmt.Println("子串最后一次出现的位置：", count)

	// 返回s中每个单词的首字母都改为标题格式的字符串拷贝。
	new_str := strings.Title(str)
	fmt.Println("单词首字母转大写后：", new_str)

	// 返回将所有字母都转为对应的大写版本的拷贝。
	new_str = strings.ToUpper(str)
	fmt.Println("所有字母转大写后：", new_str)

	// 返回将所有字母都转为对应的小写版本的拷贝。
	new_str = strings.ToLower(new_str)
	fmt.Println("所有字母转小写后：", new_str)

	// 返回将所有字母都转为对应的标题版本的拷贝。
	new_str = strings.ToTitle(str)
	fmt.Println("所有字母转标题版本大写后：", new_str)

	// 返回count个s串联的字符串。
	new_str = strings.Repeat("kj", 2)
	fmt.Println("拼接后的字符串：", new_str)

	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	new_str = strings.Replace(str, "s", "S", 2)
	fmt.Println("替换后的字符串：", new_str)

	// 返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
	new_str = strings.TrimSpace(str)
	fmt.Println("删除两端空白后的字符串：", new_str)

	// 返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	new_str = strings.TrimLeft(new_str, "this")
	fmt.Println("删除前端指定字符后的字符串：", new_str)

	// 返回将s后端所有cutset包含的utf-8码值都去掉的字符串。
	new_str = strings.TrimRight(new_str, "ing")
	fmt.Println("删除后端指定字符后的字符串：", new_str)

	// 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
	new_map := strings.Fields(str)
	fmt.Printf("转为切片后的：%q\n", new_map)
	fmt.Println("取切片中的值：", new_map[1])

	// 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	new_map = strings.Split(str, "is")
	fmt.Printf("转为切片后的：%q\n", new_map)
	fmt.Println("取切片中的值：", new_map[0])

	// 将一系列字符串连接为一个字符串，之间用sep来分隔。
	slice := []string{"java", "go", "php"}
	new_str = strings.Join(slice, "++")
	fmt.Println("切片转字符串：", new_str)

	// 获取字符串长度
	len := len(str)
	fmt.Println("获取字符串长度：", len)
}
