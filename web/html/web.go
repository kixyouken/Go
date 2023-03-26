package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/form", handleForm)
	// 分配变量
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/user", handleUser)
	http.ListenAndServe(":9090", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入响应
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	tmpl, err := template.ParseFiles("form/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入响应
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// 定义模板数据
	data := struct {
		Title   string
		Message string
	}{
		Title:   "My Page",
		Message: "Hello, World!",
	}

	// 解析模板文件
	tmpl, err := template.ParseFiles("data.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入响应
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	// 定义模板数据
	type Person struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// 定义二维结构体
	var persons = []Person{
		{Id: 1, Name: "张三", Age: 18},
		{Id: 2, Name: "李四", Age: 20},
		{Id: 3, Name: "王五"},
		{Id: 4, Name: "赵六"},
	}

	data := struct {
		Title   string
		Message string
		Data    interface{}
	}{
		Title:   "这是标题",
		Message: "Hello, World!",
		Data:    persons,
	}

	// 解析模板文件
	tmpl, err := template.ParseFiles("user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 渲染模板并将结果写入响应
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
