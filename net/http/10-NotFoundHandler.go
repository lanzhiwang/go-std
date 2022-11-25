package main

/*
type ServeMux struct {}
func NewServeMux() *ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)


*/

import (
	"fmt"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the people handler.")
	})
}

func main() {
	mux := http.NewServeMux()

	// Create sample handler to returns 404
	mux.Handle("/resources", http.NotFoundHandler())

	// Create sample handler that returns 200
	mux.Handle("/resources/people/", newPeopleHandler())

	log.Fatal(http.ListenAndServe(":8080", mux))
}
