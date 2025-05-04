package main

import (
	"fmt"
	"net/http"
	"os"
)

func fn(w http.ResponseWriter, r *http.Request) {
	bytes, _ := os.ReadFile("hello.txt")
	fmt.Fprintln(w, string(bytes))
}

func main() {
	http.HandleFunc("/hello", fn)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
