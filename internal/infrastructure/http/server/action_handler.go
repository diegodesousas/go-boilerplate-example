package server

import (
	"context"
	"net/http"
)

type Handler func(w http.ResponseWriter, req *http.Request) error

func ErrorHandler(handler Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := handler(w, req)

		if err != nil {
			handleError(req.Context(), w, err)
		}
	})
}

func handleError(ctx context.Context, w http.ResponseWriter, err error) {}
