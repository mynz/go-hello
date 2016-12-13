package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("hello")

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println("path", path)
		return nil
	})

}
