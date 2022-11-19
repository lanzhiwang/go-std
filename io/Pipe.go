package main

/*
func Pipe() (*PipeReader, *PipeWriter)
*/

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

/*
$ go run Pipe.go
some io.Reader stream to be read
*/
