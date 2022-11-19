package main

/*
type Time struct {
	// contains filtered or unexported fields
}
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func Now() Time
func Parse(layout, value string) (Time, error)
func ParseInLocation(layout, value string, loc *Location) (Time, error)
func Unix(sec int64, nsec int64) Time
func UnixMicro(usec int64) Time
func UnixMilli(msec int64) Time
func (t Time) Add(d Duration) Time
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) After(u Time) bool
func (t Time) AppendFormat(b []byte, layout string) []byte
func (t Time) Before(u Time) bool
func (t Time) Clock() (hour, min, sec int)
func (t Time) Date() (year int, month Month, day int)
func (t Time) Day() int
func (t Time) Equal(u Time) bool
func (t Time) Format(layout string) string
func (t Time) GoString() string
func (t *Time) GobDecode(data []byte) error
func (t Time) GobEncode() ([]byte, error)
func (t Time) Hour() int
func (t Time) ISOWeek() (year, week int)
func (t Time) In(loc *Location) Time
func (t Time) IsDST() bool
func (t Time) IsZero() bool
func (t Time) Local() Time
func (t Time) Location() *Location
func (t Time) MarshalBinary() ([]byte, error)
func (t Time) MarshalJSON() ([]byte, error)
func (t Time) MarshalText() ([]byte, error)
func (t Time) Minute() int
func (t Time) Month() Month
func (t Time) Nanosecond() int
func (t Time) Round(d Duration) Time
func (t Time) Second() int
func (t Time) String() string
func (t Time) Sub(u Time) Duration
func (t Time) Truncate(d Duration) Time
func (t Time) UTC() Time
func (t Time) Unix() int64
func (t Time) UnixMicro() int64
func (t Time) UnixMilli() int64
func (t Time) UnixNano() int64
func (t *Time) UnmarshalBinary(data []byte) error
func (t *Time) UnmarshalJSON(data []byte) error
func (t *Time) UnmarshalText(data []byte) error
func (t Time) Weekday() Weekday
func (t Time) Year() int
func (t Time) YearDay() int
func (t Time) Zone() (name string, offset int)
func (t Time) ZoneBounds() (start, end Time)
*/

import (
	"fmt"
	"time"
)

func expensiveCall() {}

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

/*
$ go run Duration.go
The call took 121ns to run.
*/
