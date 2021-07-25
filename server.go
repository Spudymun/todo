package todo

import (
	"context"
	"net/http"
	"time"
)

// Абстракция над http.Server
type Server struct {
	httpServer *http.Server
}

// Инкапсуляция значения для полей MaxHeaderBytes, ReadTimeout, WriteTimeout
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
