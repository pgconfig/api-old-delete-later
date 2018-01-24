package main

import (
	"encoding/json"
	"fmt"

	"github.com/pgconfig/api/categories"
	"github.com/pgconfig/api/params"
)

func main() {

	catList := categories.AllCategories

	conf := params.Input{
		PGVersion: 9.2,
		HideDoc:   true,
	}

	for i := 0; i < len(catList); i++ {
		err := catList[i].Compute(conf)
		if err != nil {
			panic(err)
		}
	}

	b, err := json.Marshal(catList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
