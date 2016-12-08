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
	// Members   []Member

	Members []interface{} `json:"members"`
}

func main() {

	g := Group{
		"the team",

		[]interface{}{
			"hoge",
			Member{"mynz"},
		},

		/*
		 * []Member{
		 *     Member{"mynz"},
		 *     Member{"jagajin"},
		 * },
		 */
	}

	doc := struct {
		Group Group
	}{
		g,
	}

	bs, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

}
