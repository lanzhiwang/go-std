const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)

var EOF = errors.New("EOF")
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
var ErrShortBuffer = errors.New("short buffer")
var ErrShortWrite = errors.New("short write")
var ErrUnexpectedEOF = errors.New("unexpected EOF")





func Copy(dst Writer, src Reader) (written int64, err error)
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
func Pipe() (*PipeReader, *PipeWriter)
func ReadAll(r Reader) ([]byte, error)
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
func WriteString(w Writer, s string) (n int, err error)


type ByteReader interface {
	ReadByte() (byte, error)
}


type ByteScanner interface {
	ByteReader
	UnreadByte() error
}


type ByteWriter interface {
	WriteByte(c byte) error
}


type Closer interface {
	Close() error
}


type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}
func (l *LimitedReader) Read(p []byte) (n int, err error)


type PipeReader struct {
	// contains filtered or unexported fields
}
func (r *PipeReader) Close() error
func (r *PipeReader) CloseWithError(err error) error
func (r *PipeReader) Read(data []byte) (n int, err error)


type PipeWriter struct {
	// contains filtered or unexported fields
}
func (w *PipeWriter) Close() error
func (w *PipeWriter) CloseWithError(err error) error
func (w *PipeWriter) Write(data []byte) (n int, err error)


type ReadCloser interface {
	Reader
	Closer
}
func NopCloser(r Reader) ReadCloser


type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}


type ReadSeeker interface {
	Reader
	Seeker
}


type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}


type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}


type ReadWriter interface {
	Reader
	Writer
}


type Reader interface {
	Read(p []byte) (n int, err error)
}
func LimitReader(r Reader, n int64) Reader
func MultiReader(readers ...Reader) Reader
func TeeReader(r Reader, w Writer) Reader


type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}


type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}


type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}


type RuneScanner interface {
	RuneReader
	UnreadRune() error
}



type SectionReader struct {
	// contains filtered or unexported fields
}
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
func (s *SectionReader) Read(p []byte) (n int, err error)
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
func (s *SectionReader) Size() int64


type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}


type StringWriter interface {
	WriteString(s string) (n int, err error)
}


type WriteCloser interface {
	Writer
	Closer
}


type WriteSeeker interface {
	Writer
	Seeker
}


type Writer interface {
	Write(p []byte) (n int, err error)
}
var Discard Writer = discard{}
func MultiWriter(writers ...Writer) Writer


type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}


type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}

