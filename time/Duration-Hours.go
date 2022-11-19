package main

/*
type Duration int64
func ParseDuration(s string) (Duration, error)
func Since(t Time) Duration
func Until(t Time) Duration
func (d Duration) Abs() Duration
func (d Duration) Hours() float64
func (d Duration) Microseconds() int64
func (d Duration) Milliseconds() int64
func (d Duration) Minutes() float64
func (d Duration) Nanoseconds() int64
func (d Duration) Round(m Duration) Duration
func (d Duration) Seconds() float64
func (d Duration) String() string
func (d Duration) Truncate(m Duration) Duration
*/

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
}

/*
$ go run Duration-Hours.go
I've got 4.5 hours of work left.
*/
