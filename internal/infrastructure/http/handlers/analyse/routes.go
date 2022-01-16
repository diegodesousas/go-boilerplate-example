package analyse

import (
	"net/http"

	"github.com/diegodesousas/go-boilerplate-example/internal/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/http/server"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/queue"
)

func Routes(factory provider.Factory, publisher queue.Publisher) server.Config {
	return func(s *server.Server) {
		s.Route(server.Route{
			Method:  http.MethodPost,
			Path:    "/analyse",
			Handler: server.ErrorHandler(Analyse(factory, publisher)),
		})
	}
}
