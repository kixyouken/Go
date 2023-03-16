package main

import (
	"encoding/json"
	"fmt"
)

// 使用接口形式可以追加元素
type Student map[string]interface{}
type Address map[string]interface{}

func main() {
	student := Student{
		"Name": "张三",
		"Age":  18,
		"Address": Address{
			"Road":     "renmin south road",
			"Street":   "123 street",
			"City":     "cs",
			"Province": "hn",
			"Country":  "CN",
			"AddText":  "这是追加的信息",
		},
		"Year":       2022, // 新增学年
		"GraduateAt": 2026, // 毕业年份
	}

	// student_info, err := json.Marshal(student)
	// 格式化json
	student_info, err := json.MarshalIndent(student, "", "    ")
	if err == nil {
		fmt.Println(string(student_info))
	} else {
		fmt.Println(err)
	}
}
