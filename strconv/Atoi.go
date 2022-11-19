package main

/*
func Atoi(s string) (int, error)
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

}

/*
$ go run Atoi.go
int, 10
*/
