package analyse

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/diegodesousas/go-boilerplate-example/internal/application/analyse"
	"github.com/diegodesousas/go-boilerplate-example/internal/domain/analysis"
	"github.com/diegodesousas/go-boilerplate-example/internal/domain/provider"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/http/server"
	"github.com/diegodesousas/go-boilerplate-example/internal/infrastructure/queue"
)

var (
	ErrAnalyseTypeNotPresent = errors.New("param X-Analyse-Type not present in header")
	ErrAnalyseTypeInvalid    = errors.New("param X-Analyse-Type in header is invalid")
)

func getAnalyseType(req *http.Request) (analysis.Type, error) {
	analyseType, ok := req.Header["X-Analyse-Type"]
	if !ok {
		return "", ErrAnalyseTypeNotPresent
	}

	t := analysis.Type(analyseType[0])

	if t != analysis.Sync && t != analysis.Async {
		return "", ErrAnalyseTypeInvalid
	}

	return t, nil
}

func sync(ctx context.Context, factory provider.Factory, a analysis.Analysis) (queue.Message, []byte, error) {
	r, err := analyse.New(factory)(ctx, a)
	if err != nil {
		return queue.Message{}, nil, err
	}

	result := analysis.Result{
		Result:   r,
		Analysis: a,
	}

	data, err := json.Marshal(result)
	if err != nil {
		return queue.Message{}, nil, err
	}

	message := queue.Message{
		Content:    data,
		RoutingKey: "result",
		Exchange:   "analysis",
	}

	return message, data, nil
}

func async(ctx context.Context, a analysis.Analysis) (queue.Message, []byte, error) {
	if err := analyse.Validate(ctx, a); err != nil {
		return queue.Message{}, nil, err
	}

	data, err := json.Marshal(a)
	if err != nil {
		return queue.Message{}, nil, err
	}

	message := queue.Message{
		Content:    data,
		RoutingKey: "analysis",
		Exchange:   "analysis",
	}

	result := analysis.Result{
		Result: provider.Result{
			Status: "analysis_sent",
		},
		Analysis: a,
	}

	data, err = json.Marshal(result)
	if err != nil {
		return queue.Message{}, nil, err
	}

	return message, data, nil
}

func Analyse(factory provider.Factory, publisher queue.Publisher) server.Handler {
	return func(w http.ResponseWriter, req *http.Request) error {
		ctx := req.Context()

		a := analysis.New()

		if err := json.NewDecoder(req.Body).Decode(&a); err != nil {
			return err
		}

		analyseType, err := getAnalyseType(req)
		if err != nil {
			return err
		}

		var message queue.Message
		var data []byte

		switch analyseType {
		case analysis.Sync:
			message, data, err = sync(ctx, factory, a)
			if err != nil {
				return err
			}
			break
		case analysis.Async:
			message, data, err = async(ctx, a)
			if err != nil {
				return err
			}
			break
		}

		if err := publisher.Publish(ctx, message); err != nil {
			return err
		}

		if _, err = w.Write(data); err != nil {
			return err
		}

		return nil
	}
}
