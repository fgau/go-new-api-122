package main

import (
	"go-new-api-122/middleware"
	"go-new-api-122/router"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
		//or db connection etc.
	}
}

func (s *APIServer) Run() error {
	router := router.SetupRoutes()

	server := http.Server{
		Addr:    s.addr,
		Handler: middleware.RequestLoggerMiddleware(router),
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}
