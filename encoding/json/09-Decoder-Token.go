package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}

/*
$ go run 09-Decoder-Token.go
json.Delim: { (more)
string: Message (more)
string: Hello (more)
string: Array (more)
json.Delim: [ (more)
float64: 1 (more)
float64: 2 (more)
float64: 3
json.Delim: ] (more)
string: Null (more)
<nil>: <nil> (more)
string: Number (more)
float64: 1.234
json.Delim: }
*/
