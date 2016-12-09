package main

import (
	"fmt"
	"net/http"
)

type MyHandle struct {
	count int
}

func (h *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.count++
	fmt.Fprintf(w, "MyHandle: count: %d", h.count)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
	fmt.Fprintf(w, "w: %v, r: %v\n", w, r)
}

func main() {
	fmt.Println("Hello")

	// http.HandleFunc("/", handler)
	http.Handle("/my", &MyHandle{})

	fs := http.FileServer(http.Dir("."))
	http.Handle("/file/", http.StripPrefix("/file", fs))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
