package logger

import (
	"net/http"
	"strings"

	"github.com/broemp/growBro/config"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func Init() {
	logLevel := config.Env.LogLevel
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	if strings.ToLower(logLevel) == "debug" {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		zap.L().Debug("http request",
			zap.String("path", r.URL.Path),
			zap.String("method", r.Method),
			zap.String("User-Agent", r.UserAgent()),
			zap.String("requestID", r.Context().Value(middleware.RequestIDKey).(string)),
			zap.String("realIP", r.Header.Get("X-Real-IP")),
			zap.String("IP", r.RemoteAddr))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
