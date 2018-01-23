package main

import (
	"encoding/json"
	"fmt"

	"github.com/pgconfig/api/params"
)

func main() {
	param := params.CheckPointSegments
	err := param.Compute(params.Input{PGVersion: 9.2})

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
