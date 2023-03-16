package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/json", handleJson)
	http.ListenAndServe(":9090", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	tmpl, err := template.ParseFiles("form.html")
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

func handlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	age := r.Form.Get("age")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Age: %s\n", age)
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	age := r.Form.Get("age")
	ageInt, _ := strconv.Atoi(age)
	person := Person{
		Name: r.Form.Get("name"),
		Age:  ageInt,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
