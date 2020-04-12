package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// 简单web页面开发

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit" />
		</form>
	</body></html>
`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>") // 输出内容到客户端第1种方式
	fmt.Fprintf(w, "login success")            // 输出内容到客户端第2种方式
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		println("param: ", request.FormValue("name"))
		io.WriteString(w, form)
	case "POST":
		err := request.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		fmt.Println("variable in:", request.Form.Get("in")) // 表单获取方式1

		out := strings.Join([]string{"Your commit is: ", request.FormValue("in")}, "") // 表单获取方式2
		io.WriteString(w, out)
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
