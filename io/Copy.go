package main

/*
func Copy(dst Writer, src Reader) (written int64, err error)
*/

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

/*
$ go run Copy.go
some io.Reader stream to be read
*/
