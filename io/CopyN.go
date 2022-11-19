package main

/*
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
*/

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}

}

/*
$ go run CopyN.go
some
*/
