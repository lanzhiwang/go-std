package main

/*
func LastIndex(s, substr string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}

/*
$ go run LastIndex.go
0
3
-1
*/
