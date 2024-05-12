package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port              = ":8080"
	pathHello         = "/hello"
	pathForm          = "/form"
	errorNotFoundMsg  = "404 not found"
	errorMethodMsg    = "Method is not supported"
	successPostMsg    = "POST request successful"
	startingServerMsg = "Starting server at port 8080"
)

// handleForm 处理表单提交的 POST 请求。
func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	response := fmt.Sprintf("%s\nName = %s\nAddress = %s", successPostMsg, name, address)
	fmt.Fprintln(w, response)
}

// handleHello 处理对 /hello 路径的 GET 请求。
func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != pathHello {
		http.Error(w, errorNotFoundMsg, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, errorMethodMsg, http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "hello!")
}

// setupRoutes 配置服务器的路由。
func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc(pathForm, handleForm)
	http.HandleFunc(pathHello, handleHello)
}

func main() {
	setupRoutes()
	log.Println(startingServerMsg)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
