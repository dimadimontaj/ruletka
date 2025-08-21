package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chim "github.com/go-chi/chi/v5/middleware"
	chimw "github.com/oapi-codegen/nethttp-middleware"

	"cases/internal/deps"
	"cases/internal/generated"
)

type Server struct {
	logger  deps.Logger
	address string
}

func NewServer(
	log deps.Logger,
	address string,
) *Server {
	return &Server{
		address: address,
		logger:  log,
	}
}

func (s *Server) Run(ctx context.Context) {
	// chi router
	r := chi.NewRouter()

	// базовые middleware
	r.Use(chim.RequestID)
	r.Use(chim.RealIP)
	r.Use(chim.Recoverer)
	r.Use(chim.Timeout(60 * time.Second))

	// техэндпоинты вне OpenAPI
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// грузим swagger для валидатора (код сгенерен oapi-codegen’ом)
	swagger, err := generated.GetSwagger()
	if err != nil {
		s.logger.Error(ctx, err)
		return
	}

	swagger.Servers = nil

	// группа API под валидацией OpenAPI
	r.Group(func(r chi.Router) {
		// валидация входящих запросов по спеки (params/body)
		r.Use(chimw.OapiRequestValidator(swagger))

		// h := generated.NewStrictHandler(s, nil) // можно прокинуть StrictHTTPServerOptions
		// generated.RegisterStrictHandlers(r, h)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    s.address,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(ctx, err)
	}
}
