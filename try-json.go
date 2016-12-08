package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name string `json:"name"`
}

type Group struct {
	GroupName string
	Members   []Member
}

func main() {

	g := Group{
		"the team",
		[]Member{
			Member{"mynz"},
			Member{"jagajin"},
		},
	}

	doc := struct {
		Group Group
	}{
		g,
	}

	// fmt.Printf("g: %T\n", g)
	// fmt.Printf("doc: %T\n", doc)
	// fmt.Printf("doc: %T, %v", doc, doc)

	bs, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

}

/*
 * func main() {
 *     mem := Member{Name: "mynz"}
 *
 *     group := Group{Members{
 *         Member{"terimukuri"},
 *     }}
 *
 *     doc := map[string]Member{
 *         "jaga": {"jagajin"},
 *     }
 *
 *     doc["mynz"] = mem
 *     doc["am"] = Member{"morimoto"}
 *
 *     bs, err := json.Marshal(doc)
 *     if err != nil {
 *         panic(err)
 *     }
 *
 *     // fmt.Println(len(bs))
 *     // fmt.Println(bs)
 *     fmt.Println(string(bs))
 *
 * }
 */
