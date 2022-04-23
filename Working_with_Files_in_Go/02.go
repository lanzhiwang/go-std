package main

// Truncate a File

import (
	"log"
	"os"
)

func main() {
    // Truncate a file to 100 bytes. If file
    // is less than 100 bytes the original contents will remain
    // at the beginning, and the rest of the space is
    // filled will null bytes. If it is over 100 bytes,
    // Everything past 100 bytes will be lost. Either way
    // we will end up with exactly 100 bytes.
    // Pass in 0 to truncate to a completely empty file
    // 将文件截断为 100 字节。 如果文件小于 100 字节，则原始内容将保留在开头，其余空间将被填充为空字节。
    // 如果超过 100 字节，超过 100 字节的所有内容都将丢失。
    // 无论哪种方式，我们最终都会得到正好 100 个字节。 传入 0 以截断为完全空的文件

    // func Truncate(name string, size int64) error
    err := os.Truncate("test.txt", 100)
    if err != nil {
        log.Fatal(err)
    }
}