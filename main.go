package main

import (
	"encoding/json"
	"fmt"

	"github.com/pgconfig/api/params"
)

func main() {

	// catList := categories.MemoryCategory

	conf := params.Input{
		PGVersion: 9.2,
		// HideDoc:   true,
	}
	// err := catList.Compute(conf)

	parm := params.CheckPointSegments
	err := parm.Compute(conf)

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(parm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
