package server

import (
	"embed"
	"net/http"

	"github.com/broemp/cannaBro/config"
	"github.com/broemp/cannaBro/server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func Start(public embed.FS) {
	zap.L().Info("Starting Server",
		zap.Uint16("port", config.Env.Port),
		zap.String("url", config.Env.URL),
		zap.String("env", config.Env.LogLevel))

	r := chi.NewMux()
	r.Use(middleware.CleanPath)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "text/html", "text/css"))
	r.Use(handler.WithUser)
	r.NotFound(errorHandler(handler.HandleNotFound))

	r.Handle("/public/*", http.StripPrefix("/", http.FileServer(http.FS(public))))

	r.Get("/", errorHandler(handler.HandleHomeIndex))
	r.Get("/login", errorHandler(handler.HandleLogin))
	r.Post("/login", errorHandler(handler.HandleLoginForm))
	r.Post("/logout", errorHandler(handler.HandleLogout))

	auth := r.Group(func(group chi.Router) {
		group.Use(handler.WithAuth)
	})

	auth.Get("/diary/new", errorHandler(handler.HandleNewDiaryEntrie))

	http.ListenAndServe(":3000", r)
}

func errorHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			zap.L().Error("internal server error", zap.Error(err), zap.String("path", r.URL.Path))
		}
	}
}
