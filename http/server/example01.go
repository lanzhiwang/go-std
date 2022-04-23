package main

import (
	"net/http"
)

// type ResponseWriter interface {
// 	Header() Header
// 	Write([]byte) (int, error)
// 	WriteHeader(statusCode int)
// }

// type Request struct {
// 	Method string
// 	URL *url.URL
// 	Proto      string // "HTTP/1.0"
// 	ProtoMajor int    // 1
// 	ProtoMinor int    // 0
// 	Header Header
// 	Body io.ReadCloser
// 	GetBody func() (io.ReadCloser, error)
// 	ContentLength int64
// 	TransferEncoding []string
// 	Close bool
// 	Host string
// 	Form url.Values
// 	PostForm url.Values
// 	MultipartForm *multipart.Form
// 	Trailer Header
// 	RemoteAddr string
// 	RequestURI string
// 	TLS *tls.ConnectionState
// 	Cancel <-chan struct{}
// 	Response *Response
// }
// func NewRequest(method, url string, body io.Reader) (*Request, error)
// func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
// func ReadRequest(b *bufio.Reader) (*Request, error)
// func (r *Request) AddCookie(c *Cookie)
// func (r *Request) BasicAuth() (username, password string, ok bool)
// func (r *Request) Clone(ctx context.Context) *Request
// func (r *Request) Context() context.Context
// func (r *Request) Cookie(name string) (*Cookie, error)
// func (r *Request) Cookies() []*Cookie
// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
// func (r *Request) FormValue(key string) string
// func (r *Request) MultipartReader() (*multipart.Reader, error)
// func (r *Request) ParseForm() error
// func (r *Request) ParseMultipartForm(maxMemory int64) error
// func (r *Request) PostFormValue(key string) string
// func (r *Request) ProtoAtLeast(major, minor int) bool
// func (r *Request) Referer() string
// func (r *Request) SetBasicAuth(username, password string)
// func (r *Request) UserAgent() string
// func (r *Request) WithContext(ctx context.Context) *Request
// func (r *Request) Write(w io.Writer) error
// func (r *Request) WriteProxy(w io.Writer) error


func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", helloWorld)

	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8000", nil)
}
