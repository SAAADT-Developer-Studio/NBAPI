package server

import (
	"NBAPI/internal/config"
	"NBAPI/internal/database"
	"NBAPI/internal/modules/player"
	"NBAPI/internal/modules/team"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	httprateredis "github.com/go-chi/httprate-redis"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	// TODO: gracefull shutdown

	r.Use(httprate.Limit(
		100,
		time.Minute,
		httprate.WithKeyByIP(),
		httprateredis.WithRedisLimitCounter(&httprateredis.Config{
			Host: config.Config.RedisHost, Port: 6379,
		}),
	))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Compress(5))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	if config.Config.AppEnv == "local" {
		r.Mount("/debug", middleware.Profiler())
	}

	r.Get("/", s.HelloWorldHandler)
	r.Route("/players", player.Router)
	r.Route("/teams", team.Router)

	r.Get("/health", s.healthHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(database.Health())
	_, _ = w.Write(jsonResp)
}
