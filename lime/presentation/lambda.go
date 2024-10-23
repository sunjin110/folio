package presentation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
)

func GetLambdaHandler() (lambdaHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), err error) {
	httpHandler, err := NewHttpHandler()
	if err != nil {
		return nil, fmt.Errorf("failed NewHttpHandler. err: %w", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", httpHandler.Home).Methods(http.MethodGet)
	r.HandleFunc("/hello", httpHandler.Hello).Methods(http.MethodGet)
	r.PathPrefix("/line").Path("/webhook").HandlerFunc(httpHandler.LineWebhook).Methods(http.MethodPost)
	return httpadapter.New(r).ProxyWithContext, nil
}
