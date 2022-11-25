```go

func CanonicalHeaderKey(s string) string
func DetectContentType(data []byte) string
func Error(w ResponseWriter, error string, code int)
func Handle(pattern string, handler Handler)
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func ListenAndServe(addr string, handler Handler) error
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
func NotFound(w ResponseWriter, r *Request)
func ParseHTTPVersion(vers string) (major, minor int, ok bool)
func ParseTime(text string) (t time.Time, err error)
func ProxyFromEnvironment(req *Request) (*url.URL, error)
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
func Redirect(w ResponseWriter, r *Request, url string, code int)
func Serve(l net.Listener, handler Handler) error
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, ...)
func ServeFile(w ResponseWriter, r *Request, name string)
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
func SetCookie(w ResponseWriter, cookie *Cookie)
func StatusText(code int) string


type Client struct {
	Transport RoundTripper
	CheckRedirect func(req *Request, via []*Request) error
	Jar CookieJar
	Timeout time.Duration
}
func (c *Client) CloseIdleConnections()
func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)


type ConnState int
func (c ConnState) String() string


type Cookie struct {
	Name  string
	Value string
	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
func (c *Cookie) String() string
func (c *Cookie) Valid() error


type CookieJar interface {
	SetCookies(u *url.URL, cookies []*Cookie)
	Cookies(u *url.URL) []*Cookie
}


type Dir string
func (d Dir) Open(name string) (File, error)


type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}


type FileSystem interface {
	Open(name string) (File, error)
}
func FS(fsys fs.FS) FileSystem


type Flusher interface {
	Flush()
}


type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
func AllowQuerySemicolons(h Handler) Handler
func FileServer(root FileSystem) Handler
func MaxBytesHandler(h Handler, n int64) Handler
func NotFoundHandler() Handler
func RedirectHandler(url string, code int) Handler
func StripPrefix(prefix string, h Handler) Handler
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler


type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)


type Header map[string][]string
func (h Header) Add(key, value string)
func (h Header) Clone() Header
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Values(key string) []string
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error


type Hijacker interface {
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}


type MaxBytesError struct {
	Limit int64
}
func (e *MaxBytesError) Error() string


type PushOptions struct {
	Method string
	Header Header
}


type Pusher interface {
	Push(target string, opts *PushOptions) error
}


type Request struct {
	Method string
	URL *url.URL
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	Header Header
	Body io.ReadCloser
	GetBody func() (io.ReadCloser, error)
	ContentLength int64
	TransferEncoding []string
	Close bool
	Host string
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form
	Trailer Header
	RemoteAddr string
	RequestURI string
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
}
func NewRequest(method, url string, body io.Reader) (*Request, error)
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
func ReadRequest(b *bufio.Reader) (*Request, error)
func (r *Request) AddCookie(c *Cookie)
func (r *Request) BasicAuth() (username, password string, ok bool)
func (r *Request) Clone(ctx context.Context) *Request
func (r *Request) Context() context.Context
func (r *Request) Cookie(name string) (*Cookie, error)
func (r *Request) Cookies() []*Cookie
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func (r *Request) FormValue(key string) string
func (r *Request) MultipartReader() (*multipart.Reader, error)
func (r *Request) ParseForm() error
func (r *Request) ParseMultipartForm(maxMemory int64) error
func (r *Request) PostFormValue(key string) string
func (r *Request) ProtoAtLeast(major, minor int) bool
func (r *Request) Referer() string
func (r *Request) SetBasicAuth(username, password string)
func (r *Request) UserAgent() string
func (r *Request) WithContext(ctx context.Context) *Request
func (r *Request) Write(w io.Writer) error
func (r *Request) WriteProxy(w io.Writer) error


type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header Header
	Body io.ReadCloser
	ContentLength int64
	TransferEncoding []string
	Close bool
	Uncompressed bool
	Trailer Header
	Request *Request
	TLS *tls.ConnectionState
}
func Get(url string) (resp *Response, err error)
func Head(url string) (resp *Response, err error)
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
func (r *Response) Cookies() []*Cookie
func (r *Response) Location() (*url.URL, error)
func (r *Response) ProtoAtLeast(major, minor int) bool
func (r *Response) Write(w io.Writer) error


type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}


type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}
func NewFileTransport(fs FileSystem) RoundTripper


type SameSite int


type ServeMux struct {}
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)


type Server struct {
	Addr string
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	TLSConfig *tls.Config
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	MaxHeaderBytes int
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
	BaseContext func(net.Listener) context.Context
	ConnContext func(ctx context.Context, c net.Conn) context.Context
}
func (srv *Server) Close() error
func (srv *Server) ListenAndServe() error
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
func (srv *Server) RegisterOnShutdown(f func())
func (srv *Server) Serve(l net.Listener) error
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
func (srv *Server) SetKeepAlivesEnabled(v bool)
func (srv *Server) Shutdown(ctx context.Context) error


type Transport struct {
	Proxy func(*Request) (*url.URL, error)
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)
	Dial func(network, addr string) (net.Conn, error)
	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)
	DialTLS func(network, addr string) (net.Conn, error)
	TLSClientConfig *tls.Config
	TLSHandshakeTimeout time.Duration
	DisableKeepAlives bool
	DisableCompression bool
	MaxIdleConns int
	MaxIdleConnsPerHost int
	MaxConnsPerHost int
	IdleConnTimeout time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper
	ProxyConnectHeader Header
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)
	MaxResponseHeaderBytes int64
	WriteBufferSize int
	ReadBufferSize int
	ForceAttemptHTTP2 bool
}
func (t *Transport) Clone() *Transport
func (t *Transport) CloseIdleConnections()
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
func (t *Transport) RoundTrip(req *Request) (*Response, error)


```

参考

* https://gowebexamples.com/

