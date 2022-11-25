package main

/*
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





*/

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var srv http.Server

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
