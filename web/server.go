package web

import (
	"net/http"
)

// Server Represents a wrapper around a net/http server
type Server struct {
	srv *http.Server
}

// Start the server
func (server *Server) Start() error {
	return server.srv.ListenAndServe()
}

// NewServer creates a new Server instance
func NewServer(addr string) *Server {
	server := new(Server)
	server.srv = new(http.Server)
	server.srv.Addr = addr
	return server
}
