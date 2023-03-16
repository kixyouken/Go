package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Road     string
	Street   string
	City     string
	Province string
	Country  string
}

func main() {
	student := Student{
		Name: "张三",
		Age:  18,
		Address: Address{
			Road:     "renmin south road",
			Street:   "123 street",
			City:     "cs",
			Province: "hn",
			Country:  "CN",
		},
	}

	// student_info, err := json.Marshal(student)
	// 格式化json
	student_info, err := json.MarshalIndent(student, "", "    ")
	if err == nil {
		fmt.Println(string(student_info))
	} else {
		fmt.Println(err)
		return
	}
}
