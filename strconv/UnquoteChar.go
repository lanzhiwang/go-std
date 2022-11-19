package main

/*
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
*/

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)

}

/*
$ go run UnquoteChar.go
value: "
multibyte: false
tail: Fran & Freddie's Diner\"
*/
