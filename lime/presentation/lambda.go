package presentation

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
)

func GetHandler() (lambdaHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), err error) {
	r := mux.NewRouter()

	slog.Info("routerは作った")
	r.Get("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.Info("======= / がリクエストされた")

		slog.Info("base")
		w.Write([]byte(`{"hello": "lime"}`))
		w.WriteHeader(http.StatusOK)
		slog.Info("====== / の処理終わり")
	})
	r.Get("/hello").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("hello lime")
		w.Write([]byte(`{"hello": "lime"}`))
		w.WriteHeader(http.StatusOK)
		slog.Info("======== /helloの処理終わり")
	})

	slog.Info("ここまできた")
	return httpadapter.New(r).ProxyWithContext, nil
}
