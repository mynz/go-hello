package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "foo"
	queue <- "bar"
	queue <- "baz"
	queue <- "hoge"

	cancel()
	wg.Wait()

	fmt.Println("bye")
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exits")
			wg.Done()
			return
		case url := <-queue:
			fmt.Println("fetch URL", url)
		}
	}
}
