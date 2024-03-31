package main

import (
	"go-new-api-122/router"
	"log"
	"net/http"
	"time"
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
		Handler: RequestLoggerMiddleware(router),
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("method %s, path: %s in %v", r.Method, r.URL.Path, duration)
	}
}
