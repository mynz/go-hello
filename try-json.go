package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
}

func main() {
	mes := Message{Name: "mynz"}

	doc := map[string]Message{
		"jaga": {"jagajin"},
	}

	doc["mynz"] = mes
	doc["am"] = Message{"morimoto"}

	bs, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(bs))
	// fmt.Println(bs)
	fmt.Println(string(bs))

}
