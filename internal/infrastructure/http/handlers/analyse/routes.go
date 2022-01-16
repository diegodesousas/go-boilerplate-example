package analyse

import (
	"net/http"

	"github.com/diegodesousas/go-boilerplate-example/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/http/server"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/queue"
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
