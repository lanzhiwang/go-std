package main

/*
func Valid(data []byte) bool
*/

import (
	"encoding/json"
	"fmt"
)

func main() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}

/*
$ go run 06-Valid.go
true false
*/
