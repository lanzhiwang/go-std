package main

/*
func HTMLEscape(dst *bytes.Buffer, src []byte)
*/

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}

/*
$ go run 05-HTMLEscape.go
{"Name":"\u003cb\u003eHTML content\u003c/b\u003e"}
*/
