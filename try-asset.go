package main

import (
	// "bytes"
	"fmt"
)

//go:generate go-bindata -prefix data ./data

func main() {

	var err error

	/*
	 * a, err := Asset("hoge.txt")
	 * if err != nil {
	 *     panic(err)
	 * }
	 */

	a := MustAsset("hoge.txt")

	fmt.Printf("asset: %T, %v", a, a)
	fmt.Println("asset", a, err)

	fmt.Println("str", string(a[:]))
	fmt.Println("b", []byte(string(a)))

}
