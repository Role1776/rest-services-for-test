package app

import (
	"net/http"
	"time"
)

type server struct {
	httpserver *http.Server
}

func (s *server) Run(port string, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
	}

	return s.httpserver.ListenAndServe()
}

func InitServer() *server {
	return &server{}
}
