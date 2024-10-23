package presentation

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
)

func GetHandler() (lambdaHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), err error) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "home"}`))
	}).Methods(http.MethodGet)
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "lime"}`))
	}).Methods(http.MethodGet)
	return httpadapter.New(r).ProxyWithContext, nil
}
