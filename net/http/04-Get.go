package main

/*
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

*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
