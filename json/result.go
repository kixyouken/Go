package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    interface{}
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 定义二维结构体
	var persons = []Person{
		{Id: 1, Name: "张三", Age: 18},
		{Id: 2, Name: "李四", Age: 20},
		{Id: 3, Name: "王五"},
		{Id: 4, Name: "赵六"},
	}
	result := Result{
		Code:    200,
		Message: "success",
		Data:    persons,
	}
	// result_info, err := json.Marshal(result)
	// 格式化json
	result_info, err := json.MarshalIndent(result, "", "    ")
	if err == nil {
		fmt.Println(string(result_info))
	} else {
		fmt.Println(err)
		return
	}
}
