package main

/*
func ReadAll(r Reader) ([]byte, error)
*/

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}

/*
$ go run ReadAll.go
Go is a general-purpose language designed with systems programming in mind.
*/
