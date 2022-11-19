package main

/*
func Unmarshal(data []byte, v any) error
*/

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

/*
$ go run 02-Unmarshal.go
[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
*/
