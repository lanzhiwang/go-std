var Discard io.Writer = io.Discard

func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]fs.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, pattern string) (name string, err error)
func TempFile(dir, pattern string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm fs.FileMode) error
