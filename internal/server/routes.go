package server

import (
	"NBAPI/internal/config"
	"NBAPI/internal/database"
	"NBAPI/internal/modules/player"
	"NBAPI/internal/modules/team"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	httprateredis "github.com/go-chi/httprate-redis"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

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
	api := humachi.New(r, huma.DefaultConfig("My API", "1.0.0"))

	huma.Get(api, "/", s.HelloWorldHandler)
	r.Route("/players", player.Router)
	team.RegisterRoutes(api)

	r.Get("/health", s.healthHandler)

	r.Get("/docs", s.DocsHandler)

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

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func (s *Server) HelloWorldHandler(ctx context.Context, input *struct{}) (*GreetingOutput, error) {
	resp := &GreetingOutput{}
	resp.Body.Message = "Hello world"
	return resp, nil
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(database.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) DocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!doctype html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      data-url="/openapi.json"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`))
}
