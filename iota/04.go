package main

import "fmt"

type Order int

const (
	// LSB means Least Significant Bits first, as used in the GIF file format.
	LSB Order = iota
	// MSB means Most Significant Bits first, as used in the TIFF and PDF
	// file formats.
	MSB
)

func main() {
	fmt.Println("LSB: ", LSB)
	fmt.Println("MSB: ", MSB)
}

/*
LSB:  0
MSB:  1
*/
