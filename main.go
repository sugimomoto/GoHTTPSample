package main

import (
	"html/template"
	"net/http"
)

func simpleTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("simple.html")
	if err != nil {
		panic(err)
	}

	// 単一の変数を渡すことが可能
	// {{ . }} に置換される
	t.Execute(writer, "Template")
}

func structTemplate(writer http.ResponseWriter, request *http.Request) {
	type StructSample struct {
		Title   string
		Message string
	}

	structSample := StructSample{
		Title:   "Struct Sample",
		Message: "Welcome Page!",
	}

	t, err := template.ParseFiles("struct.html")
	if err != nil {
		panic(err)
	}

	// 構造体を渡すこともできる。
	// .Title などの指定でアクセスが可能
	t.Execute(writer, structSample)
}

func arrayTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("array.html")
	if err != nil {
		panic(err)
	}

	arraySample := []string{
		"Orange",
		"Gray",
		"Red",
	}

	// 配列はrangeで取り出す
	// {{ range . }} に置換される
	t.Execute(writer, arraySample)
}

func main() {
	http.HandleFunc("/simple", simpleTemplate)
	http.HandleFunc("/struct", structTemplate)
	http.HandleFunc("/array", arrayTemplate)
	http.ListenAndServe(":8081", nil)
}
