package main

/*
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
*/

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

}

/*
$ go run CopyBuffer.go
first reader

second reader

*/
