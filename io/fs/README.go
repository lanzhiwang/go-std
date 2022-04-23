var (
	ErrInvalid    = errInvalid()    // "invalid argument"
	ErrPermission = errPermission() // "permission denied"
	ErrExist      = errExist()      // "file already exists"
	ErrNotExist   = errNotExist()   // "file does not exist"
	ErrClosed     = errClosed()     // "file already closed"
)

var SkipDir = errors.New("skip this directory")

func Glob(fsys FS, pattern string) (matches []string, err error)
func ReadFile(fsys FS, name string) ([]byte, error)
func ValidPath(name string) bool
func WalkDir(fsys FS, root string, fn WalkDirFunc) error


type DirEntry interface {
	Name() string
	IsDir() bool
	Type() FileMode
	Info() (FileInfo, error)
}
func FileInfoToDirEntry(info FileInfo) DirEntry
func ReadDir(fsys FS, name string) ([]DirEntry, error)


type FS interface {
	Open(name string) (File, error)
}
func Sub(fsys FS, dir string) (FS, error)


type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}


type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() any           // underlying data source (can return nil)
}
func Stat(fsys FS, name string) (FileInfo, error)


type FileMode uint32
func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool
func (m FileMode) Perm() FileMode
func (m FileMode) String() string
func (m FileMode) Type() FileMode


type GlobFS interface {
	FS
	Glob(pattern string) ([]string, error)
}


type PathError struct {
	Op   string
	Path string
	Err  error
}
func (e *PathError) Error() string
func (e *PathError) Timeout() bool
func (e *PathError) Unwrap() error


type ReadDirFS interface {
	FS
	ReadDir(name string) ([]DirEntry, error)
}


type ReadDirFile interface {
	File
	ReadDir(n int) ([]DirEntry, error)
}


type ReadFileFS interface {
	FS
	ReadFile(name string) ([]byte, error)
}


type StatFS interface {
	FS
	Stat(name string) (FileInfo, error)
}


type SubFS interface {
	FS
	Sub(dir string) (FS, error)
}


type WalkDirFunc func(path string, d DirEntry, err error) error

