package main

import (
	"fmt"
	"runtime"
	"time"
	// "sync"
)

func hoge(d time.Duration) <-chan time.Time {
	fmt.Println("call hoge")
	return time.After(d)
}

func fuga() chan bool {
	ch := make(chan bool)
	fmt.Println("call fuga")
	return ch
}

func main() {
	fmt.Println("hello")
	// fmt.Println("NumGroutine", runtime.NumGoroutine())

	fmt.Println("NumCPU", runtime.NumCPU())

	select {
	// case c := <-time.After(time.Second):
	case c := <-hoge(time.Second):
		fmt.Println("one:", c)
	case <-fuga():
		fmt.Println("two:")
	}

	fmt.Println("finish.")
}
