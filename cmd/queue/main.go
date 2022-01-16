package main

import (
	"os"

	"github.com/diegodesousas/go-boilerplate-example/infrastructure/external/provider"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/queue/handlers/analyse"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/queue/rabbitmq/consumer"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/queue/rabbitmq/publisher"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	rabbitPublisher := publisher.New()
	providerFactory := provider.NewFactory()

	analyseConsumer := consumer.New("analyse", analyse.Analyse(providerFactory, rabbitPublisher))

	interrupt := make(chan os.Signal, 1)

	go func() {
		err := consumer.Listen(
			analyseConsumer,
		)
		if err != nil {
			return
		}
	}()

	<-interrupt
	err := consumer.Shutdown()
	if err != nil {
		return
	}
}
