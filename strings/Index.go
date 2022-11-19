package main

/*
func Index(s, substr string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}

/*
$ go run Index.go
4
-1
*/
