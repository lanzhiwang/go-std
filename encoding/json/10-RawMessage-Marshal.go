package main

/*
type RawMessage
func (m RawMessage) MarshalJSON() ([]byte, error)
func (m *RawMessage) UnmarshalJSON(data []byte) error
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

/*
$ go run 09-RawMessage-Marshal.go
{
	"header": {
		"precomputed": true
	},
	"body": "Hello Gophers!"
}
*/
