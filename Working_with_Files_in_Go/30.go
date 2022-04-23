package main

// Hashing and Checksums

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
    // Open file for reading
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create new hasher, which is a writer interface
    hasher := md5.New()
    _, err = io.Copy(hasher, file)
    if err != nil {
        log.Fatal(err)
    }

    // Hash and print. Pass nil since
    // the data is not coming in as a slice argument
    // but is coming through the writer interface
    sum := hasher.Sum(nil)
    fmt.Printf("Md5 checksum: %x\n", sum)
}