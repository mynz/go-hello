package main

import (
	"fmt"
	"os"
	// "sync"
)

func main() {
	fmt.Println("hello")

	f, err := os.Open("fuga.txt")
	defer func() { println("closing"); f.Close() }()

	println("f:", f)

	if err != nil {
		panic(err)
	}

	d, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(d)

	fmt.Println("finish.")
}
