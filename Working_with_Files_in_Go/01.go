package main

// Create Empty File

import (
	"log"
	"os"
)

var (
    newFile *os.File
    err     error
)

func main() {
    // func Create(name string) (*File, error)
    newFile, err = os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    log.Println(*newFile)
    newFile.Close()
}

/*
$ go run 01.go
2022/04/23 08:59:37 &{0xc00006c180}
2022/04/23 08:59:37 {0xc00006c180}
*/
