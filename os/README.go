const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)

const (
	SEEK_SET int = 0 // seek relative to the origin of the file
	SEEK_CUR int = 1 // seek relative to the current offset
	SEEK_END int = 2 // seek relative to the end
)

const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        = fs.ModeDir        // d: is a directory
	ModeAppend     = fs.ModeAppend     // a: append-only
	ModeExclusive  = fs.ModeExclusive  // l: exclusive use
	ModeTemporary  = fs.ModeTemporary  // T: temporary file; Plan 9 only
	ModeSymlink    = fs.ModeSymlink    // L: symbolic link
	ModeDevice     = fs.ModeDevice     // D: device file
	ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)
	ModeSocket     = fs.ModeSocket     // S: Unix domain socket
	ModeSetuid     = fs.ModeSetuid     // u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid
	ModeCharDevice = fs.ModeCharDevice // c: Unix character device, when ModeDevice is set
	ModeSticky     = fs.ModeSticky     // t: sticky
	ModeIrregular  = fs.ModeIrregular  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm // Unix permission bits, 0o777
)

const DevNull = "/dev/null"

var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)

var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

var Args []string

var ErrProcessDone = errors.New("os: process already finished")


func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
func Chtimes(name string, atime time.Time, mtime time.Time) error
func Clearenv()
func DirFS(dir string) fs.FS
func Environ() []string
func Executable() (string, error)
func Exit(code int)
func Expand(s string, mapping func(string) string) string
func ExpandEnv(s string) string
func Getegid() int
func Getenv(key string) string
func Geteuid() int
func Getgid() int
func Getgroups() ([]int, error)
func Getpagesize() int
func Getpid() int
func Getppid() int
func Getuid() int
func Getwd() (dir string, err error)
func Hostname() (name string, err error)
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPathSeparator(c uint8) bool
func IsPermission(err error) bool
func IsTimeout(err error) bool
func Lchown(name string, uid, gid int) error
func Link(oldname, newname string) error
func LookupEnv(key string) (string, bool)
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
func MkdirTemp(dir, pattern string) (string, error)
func NewSyscallError(syscall string, err error) error
func Pipe() (r *File, w *File, err error)
func ReadFile(name string) ([]byte, error)
func Readlink(name string) (string, error)
func Remove(name string) error
func RemoveAll(path string) error
func Rename(oldpath, newpath string) error
func SameFile(fi1, fi2 FileInfo) bool
func Setenv(key, value string) error
func Symlink(oldname, newname string) error
func TempDir() string
func Truncate(name string, size int64) error
func Unsetenv(key string) error
func UserCacheDir() (string, error)
func UserConfigDir() (string, error)
func UserHomeDir() (string, error)
func WriteFile(name string, data []byte, perm FileMode) error


type DirEntry = fs.DirEntry
func ReadDir(name string) ([]DirEntry, error)


type File struct {
	// contains filtered or unexported fields
}
func Create(name string) (*File, error)
func CreateTemp(dir, pattern string) (*File, error)
func NewFile(fd uintptr, name string) *File
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (f *File) Close() error
func (f *File) Fd() uintptr
func (f *File) Name() string
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) ReadDir(n int) ([]DirEntry, error)
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) SetDeadline(t time.Time) error
func (f *File) SetReadDeadline(t time.Time) error
func (f *File) SetWriteDeadline(t time.Time) error
func (f *File) Stat() (FileInfo, error)
func (f *File) Sync() error
func (f *File) SyscallConn() (syscall.RawConn, error)
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)


type FileInfo = fs.FileInfo
func Lstat(name string) (FileInfo, error)
func Stat(name string) (FileInfo, error)


type FileMode = fs.FileMode


type LinkError struct {
	Op  string
	Old string
	New string
	Err error
}
func (e *LinkError) Error() string
func (e *LinkError) Unwrap() error


type PathError = fs.PathError


type ProcAttr struct {
	Dir string
	Env []string
	Files []*File
	Sys *syscall.SysProcAttr
}


type Process struct {
	Pid int
}
func FindProcess(pid int) (*Process, error)
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
func (p *Process) Kill() error
func (p *Process) Release() error
func (p *Process) Signal(sig Signal) error
func (p *Process) Wait() (*ProcessState, error)


type ProcessState struct {
	// contains filtered or unexported fields
}
func (p *ProcessState) ExitCode() int
func (p *ProcessState) Exited() bool
func (p *ProcessState) Pid() int
func (p *ProcessState) String() string
func (p *ProcessState) Success() bool
func (p *ProcessState) Sys() any
func (p *ProcessState) SysUsage() any
func (p *ProcessState) SystemTime() time.Duration
func (p *ProcessState) UserTime() time.Duration


type Signal interface {
	String() string
	Signal() // to distinguish from other Stringers
}
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)


type SyscallError struct {
	Syscall string
	Err     error
}
func (e *SyscallError) Error() string
func (e *SyscallError) Timeout() bool
func (e *SyscallError) Unwrap() error

