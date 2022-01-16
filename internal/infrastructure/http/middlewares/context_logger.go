package middlewares

import (
	"net/http"

	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l := logger.NewLogger()
		next.ServeHTTP(w, req.WithContext(logger.NewContext(req.Context(), l)))
	})
}
