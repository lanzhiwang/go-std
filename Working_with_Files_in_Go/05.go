package main

// Delete a File

import (
	"log"
	"os"
)

func main() {
    // func Remove(name string) error
    err := os.Remove("test.txt")
    if err != nil {
        log.Fatal(err)
    }
}