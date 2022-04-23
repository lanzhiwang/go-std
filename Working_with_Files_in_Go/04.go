package main

// Rename and Move a File

import (
	"log"
	"os"
)

func main() {
    originalPath := "test.txt"
    newPath := "test2.txt"

    // func Rename(oldpath, newpath string) error
    err := os.Rename(originalPath, newPath)
    if err != nil {
        log.Fatal(err)
    }
}