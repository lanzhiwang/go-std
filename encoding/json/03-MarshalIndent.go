package main

/*
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
*/

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	b, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

/*
$ go run 03-MarshalIndent.go
{
<prefix><indent>"a": 1,
<prefix><indent>"b": 2
<prefix>}
*/
