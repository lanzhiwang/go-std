package main

/*
func Marshal(v any) ([]byte, error)
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println(b)
	fmt.Println(string(b))
}

/*
$ go run 01-Marshal.go
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
[123 34 73 68 34 58 49 44 34 78 97 109 101 34 58 34 82 101 100 115 34 44 34 67 111 108 111 114 115 34 58 91 34 67 114 105 109 115 111 110 34 44 34 82 101 100 34 44 34 82 117 98 121 34 44 34 77 97 114 111 111 110 34 93 125]
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
*/



