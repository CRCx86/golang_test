package main

import (
	"encoding/json"
	"fmt"
)

// реализация структуры

// A represents a single literal
type A struct {
	RealName string `json:"real_name"`
	RealAge  int    `json:"real_age"`
}

func (a *A) getAge() int {
	return a.RealAge
}

func main() {

	a := []A{
		{
			RealName: "A",
			RealAge:  1,
		},
		{
			RealName: "C",
			RealAge:  2,
		},
	}

	jsonBytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))

}
