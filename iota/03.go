package main

import "fmt"

type ErrorHandling int

const (
	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
	PanicOnError                         // Call panic with a descriptive error.
)

func main() {
	fmt.Println("ContinueOnError: ", ContinueOnError)
	fmt.Println("ExitOnError: ", ExitOnError)
	fmt.Println("PanicOnError: ", PanicOnError)
}

/*
ContinueOnError:  0
ExitOnError:  1
PanicOnError:  2
*/
