package analyse

import (
	"context"
	"encoding/json"

	"github.com/diegodesousas/go-boilerplate-example/application/analyse"
	"github.com/diegodesousas/go-boilerplate-example/domain/analysis"
	"github.com/diegodesousas/go-boilerplate-example/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/infrastructure/queue"
)

func Analyse(factory provider.Factory, publisher queue.Publisher) queue.Handler {
	return func(ctx context.Context, message queue.Message) error {
		a := analysis.New()

		if err := json.Unmarshal(message.Content, &a); err != nil {
			return err
		}

		r, err := analyse.New(factory)(ctx, a)
		if err != nil {
			return err
		}

		result := analysis.Result{
			Result:   r,
			Analysis: a,
		}

		data, err := json.Marshal(result)
		if err != nil {
			return err
		}

		resultMessage := queue.Message{
			Content:    data,
			RoutingKey: "result",
			Exchange:   "analysis",
		}

		if err = publisher.Publish(ctx, resultMessage); err != nil {
			return err
		}

		return nil
	}
}
