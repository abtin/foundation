package app

import (
	"log"
	"net/http"

	"foundation/pkg/auth"
	"foundation/pkg/config"

	"github.com/go-chi/chi/v5"
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

// Run starts the server
func (s Server) Run() error {
	// setup http server
	router := chi.NewRouter()
	router.Get("/public",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("open to everyone"))
		})

	authService := auth.NewGithubAuth(s.Config)
	m := authService.Middleware()
	router.With(m.Auth).Get("/private",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("only if you are logged in!"))
		})

	//setup auth routes
	authRoutes, avaRoutes := authService.Handlers()
	router.Mount("/auth", authRoutes)
	router.Mount("/avatar", avaRoutes)

	log.Printf("Listening on  %s\n", s.Config.WebServerURL())
	if err := http.ListenAndServe(s.Config.WebServerURL(), router); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}
