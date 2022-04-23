package main

// Get File Info

import (
	"fmt"
	"log"
	"os"
)

var (
    fileInfo os.FileInfo
    err      error
)

func main() {
    // Stat returns file info. It will return
    // an error if there is no file.

    // func Stat(name string) (FileInfo, error)
    fileInfo, err = os.Stat("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    /*
    type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() any           // underlying data source (can return nil)
    }
    */
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
    fmt.Println("Is Directory: ", fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n", fileInfo.Sys())
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

/*
$ go run 03.go
File name: test.txt
Size in bytes: 100
Permissions: -rw-r--r--
Last modified: 2022-04-23 09:09:02.742088368 +0800 CST
Is Directory:  false
System interface type: *syscall.Stat_t
System info: &{Dev:16777221 Mode:33188 Nlink:1 Ino:43090633 Uid:501 Gid:20 Rdev:0 Pad_cgo_0:[0 0 0 0] Atimespec:{Sec:1650676139 Nsec:386612872} Mtimespec:{Sec:1650676142 Nsec:742088368} Ctimespec:{Sec:1650676142 Nsec:742088368} Birthtimespec:{Sec:1650676139 Nsec:386612872} Size:100 Blocks:0 Blksize:4096 Flags:0 Gen:0 Lspare:0 Qspare:[0 0]}
*/
