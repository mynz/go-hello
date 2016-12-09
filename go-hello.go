package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "sync"
)

func main() {
	fmt.Println("hello")

	// url := "yahoo.co.jp"
	url := "http://garbagecollection.org"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	// bs := make([]byte)
	// bs := []byte{}
	// resp.Body.Read(bs)

	bs, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bs))

	fmt.Println("finish.")
}
