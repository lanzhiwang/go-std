package main

/*
func ToTitle(s string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the Title example.
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
}

/*
$ go run ToTitle.go
HER ROYAL HIGHNESS
LOUD NOISES
ХЛЕБ
*/
