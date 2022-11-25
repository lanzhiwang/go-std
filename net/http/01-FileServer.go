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

func ListenAndServe(addr string, handler Handler) error

*/

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}





