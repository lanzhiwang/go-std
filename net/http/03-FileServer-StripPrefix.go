package main

/*
type Dir string
func (d Dir) Open(name string) (File, error)

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

func Handle(pattern string, handler Handler)
*/

import (
	"net/http"
)

func main() {
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/tmpfiles/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
}
