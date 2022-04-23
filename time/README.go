const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)


func After(d Duration) <-chan Time
func Sleep(d Duration)
func Tick(d Duration) <-chan Time


type Duration int64
func ParseDuration(s string) (Duration, error)
func Since(t Time) Duration
func Until(t Time) Duration
func (d Duration) Hours() float64
func (d Duration) Microseconds() int64
func (d Duration) Milliseconds() int64
func (d Duration) Minutes() float64
func (d Duration) Nanoseconds() int64
func (d Duration) Round(m Duration) Duration
func (d Duration) Seconds() float64
func (d Duration) String() string
func (d Duration) Truncate(m Duration) Duration


type Location struct {
	// contains filtered or unexported fields
}
func FixedZone(name string, offset int) *Location
func LoadLocation(name string) (*Location, error)
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
func (l *Location) String() string


type Month int
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
func (m Month) String() string


type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}
func (e *ParseError) Error() string


type Ticker struct {
	C <-chan Time // The channel on which the ticks are delivered.
	// contains filtered or unexported fields
}
func NewTicker(d Duration) *Ticker
func (t *Ticker) Reset(d Duration)
func (t *Ticker) Stop()


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


type Timer struct {
	C <-chan Time
	// contains filtered or unexported fields
}
func AfterFunc(d Duration, f func()) *Timer
func NewTimer(d Duration) *Timer
func (t *Timer) Reset(d Duration) bool
func (t *Timer) Stop() bool


type Weekday int
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
func (d Weekday) String() string
