package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server представляет собой структуру для HTTP-сервера
type Server struct {
	Router *mux.Router
	Port   string
	Timeout time.Duration
}

// New создает новый экземпляр Server
func New(port string, timeout time.Duration) *Server {
	return &Server{
		Router: mux.NewRouter(),
		Port:   port,
		Timeout: timeout,
	}
}

// Run запускает HTTP-сервер
func (s *Server) Run() error {
	// Настройка таймаута
	srv := &http.Server{
		Handler: s.Router,
		Addr:    ":" + s.Port,
		WriteTimeout: s.Timeout,
		ReadTimeout:  s.Timeout,
	}

	// Запуск сервера
	return srv.ListenAndServe()
}
