package main

import (
	"context"
	"log"

	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/external/provider"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/http/handlers/analyse"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/http/server"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/queue/rabbitmq/publisher"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	ctx := context.Background()

	rabbitPublisher := publisher.New()
	providerFactory := provider.NewFactory()

	s := server.NewServer(
		analyse.Routes(providerFactory, rabbitPublisher),
	)

	err := s.ListenAndServe(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}
