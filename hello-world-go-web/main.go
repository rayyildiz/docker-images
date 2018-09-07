package main

import (
	"io"	
	"net/http"
)

func hello(w http.ResponseWriter, r* http.Request) {
	io.WriteString(w,"Hello World from Golang by @rayyildiz")
}

func main() {
	http.HandleFunc("/",hello)
	http.ListenAndServe(":8080",nil)
}