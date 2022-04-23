func Compact(dst *bytes.Buffer, src []byte) error
func HTMLEscape(dst *bytes.Buffer, src []byte)
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
func Marshal(v any) ([]byte, error)
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
func Unmarshal(data []byte, v any) error
func Valid(data []byte) bool


type Decoder struct {
	// contains filtered or unexported fields
}
func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Buffered() io.Reader
func (dec *Decoder) Decode(v any) error
func (dec *Decoder) DisallowUnknownFields()
func (dec *Decoder) InputOffset() int64
func (dec *Decoder) More() bool
func (dec *Decoder) Token() (Token, error)
func (dec *Decoder) UseNumber()


type Delim rune
func (d Delim) String() string


type Encoder struct {
	// contains filtered or unexported fields
}
func NewEncoder(w io.Writer) *Encoder
func (enc *Encoder) Encode(v any) error
func (enc *Encoder) SetEscapeHTML(on bool)
func (enc *Encoder) SetIndent(prefix, indent string)


DEPRECATED
type InvalidUTF8Error struct {
	S string // the whole string value that caused the error
}
func (e *InvalidUTF8Error) Error() string


type InvalidUnmarshalError struct {
	Type reflect.Type
}
func (e *InvalidUnmarshalError) Error() string


type Marshaler interface {
	MarshalJSON() ([]byte, error)
}


type MarshalerError struct {
	Type reflect.Type
	Err  error
	// contains filtered or unexported fields
}
func (e *MarshalerError) Error() string
func (e *MarshalerError) Unwrap() error


type Number string
func (n Number) Float64() (float64, error)
func (n Number) Int64() (int64, error)
func (n Number) String() string


type RawMessage []byte
func (m RawMessage) MarshalJSON() ([]byte, error)
func (m *RawMessage) UnmarshalJSON(data []byte) error


type SyntaxError struct {
	Offset int64 // error occurred after reading Offset bytes
	// contains filtered or unexported fields
}
func (e *SyntaxError) Error() string


type Token any


DEPRECATED
type UnmarshalFieldError struct {
	Key   string
	Type  reflect.Type
	Field reflect.StructField
}
func (e *UnmarshalFieldError) Error() string


type UnmarshalTypeError struct {
	Value  string       // description of JSON value - "bool", "array", "number -5"
	Type   reflect.Type // type of Go value it could not be assigned to
	Offset int64        // error occurred after reading Offset bytes
	Struct string       // name of the struct type containing the field
	Field  string       // the full path from root node to the field
}
func (e *UnmarshalTypeError) Error() string


type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}


type UnsupportedTypeError struct {
	Type reflect.Type
}
func (e *UnsupportedTypeError) Error() string


type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}
func (e *UnsupportedValueError) Error() string
