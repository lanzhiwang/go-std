package main

import "fmt"

type FileMode uint32

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky
	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777 // Unix permission bits
)


func main() {
	fmt.Printf("ModeDir:         %b\n", ModeDir)
	fmt.Printf("ModeAppend:      %b\n", ModeAppend)
	fmt.Printf("ModeExclusive:   %b\n", ModeExclusive)
	fmt.Printf("ModeTemporary:   %b\n", ModeTemporary)
	fmt.Printf("ModeSymlink:     %b\n", ModeSymlink)
	fmt.Printf("ModeDevice:      %b\n", ModeDevice)
	fmt.Printf("ModeNamedPipe:   %b\n", ModeNamedPipe)
	fmt.Printf("ModeSocket:      %b\n", ModeSocket)
	fmt.Printf("ModeSetuid:      %b\n", ModeSetuid)
	fmt.Printf("ModeSetgid:      %b\n", ModeSetgid)
	fmt.Printf("ModeCharDevice:  %b\n", ModeCharDevice)
	fmt.Printf("ModeSticky:      %b\n", ModeSticky)
	fmt.Printf("ModeIrregular:   %b\n", ModeIrregular)
	fmt.Printf("ModeType:        %b\n", ModeType)
	fmt.Printf("ModePerm:        %b\n", ModePerm)

}

/*
ModeDir:         10000000000000000000000000000000
ModeAppend:      1000000000000000000000000000000
ModeExclusive:   100000000000000000000000000000
ModeTemporary:   10000000000000000000000000000
ModeSymlink:     1000000000000000000000000000
ModeDevice:      100000000000000000000000000
ModeNamedPipe:   10000000000000000000000000
ModeSocket:      1000000000000000000000000
ModeSetuid:      100000000000000000000000
ModeSetgid:      10000000000000000000000
ModeCharDevice:  1000000000000000000000
ModeSticky:      100000000000000000000
ModeIrregular:   10000000000000000000
ModeType:        10001111001010000000000000000000
ModePerm:        111111111
*/
