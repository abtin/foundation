package app

import (
	"fmt"
	"log"
	"net/http"

	"foundation/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server has configuration
type Server struct {
	Config config.Config
}

// NewServer creates a new Server using the give config
func NewServer(config config.Config) Server {
	return Server{
		Config: config,
	}
}

// Run srarts the server
func (s Server) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	port := fmt.Sprintf(":%d", s.Config.Port)
	log.Printf("Listening on Port %d\n", s.Config.Port)
	if err := http.ListenAndServe(port, r);err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}
