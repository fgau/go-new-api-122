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

	midllewareChain := middleware.MiddlewareChain(
		middleware.RequestLoggerMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: midllewareChain(router),
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}
