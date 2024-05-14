package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/MWT-proger/go-keeper/configs/server/config"
	"github.com/MWT-proger/go-keeper/internal/pkg/logger"
	"github.com/MWT-proger/go-keeper/internal/server/auth"
	"github.com/MWT-proger/go-keeper/internal/server/gzip"
	"github.com/MWT-proger/go-keeper/internal/server/handlers"
)

// router() Перенаправляет запросы на необходимые хендлеры
func router(h *handlers.APIHandler, conf *config.Config) *chi.Mux {

	r := chi.NewRouter()
	r.Use(logger.RequestLogger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   conf.Cors.AllowedOrigins,
		AllowedMethods:   conf.Cors.AllowedMethods,
		AllowedHeaders:   conf.Cors.AllowedHeaders,
		AllowCredentials: conf.Cors.AllowCredentials,
		Debug:            conf.Cors.Debug,
	}))
	r.Use(gzip.GzipMiddleware)

	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/user/register", h.UserRegister)
		r.Post("/user/login", h.UserLogin)

		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return auth.AuthCookieMiddleware(next, conf)
			})
			r.Post("/user/test", func(w http.ResponseWriter, r *http.Request) {})

		})

	})

	return r
}
