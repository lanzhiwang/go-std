package main

/*
func Fields(s string) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}

/*
$ go run Fields.go
Fields are: ["foo" "bar" "baz"]
*/
