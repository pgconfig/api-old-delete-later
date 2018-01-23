package main

import (
	"encoding/json"
	"fmt"

	"github.com/pgconfig/api/categories"
	"github.com/pgconfig/api/params"
)

func main() {

	category := categories.MemoryCategory

	// param := params.CheckPointSegments
	err := category.Compute(params.Input{PGVersion: 9.2})

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
