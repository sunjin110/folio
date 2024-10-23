package presentation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
	"github.com/sunjin110/folio/lime/application"
	"github.com/sunjin110/folio/lime/config"
)

func GetHandler() (lambdaHandlerFunc func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error), err error) {

	envConfig, err := config.NewEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("failed config.NewEnvConfig. err: %w", err)
	}

	lineUsecase := application.NewLineUsecase(envConfig.Line.ChannelSecret)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "home"}`))
	}).Methods(http.MethodGet)
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "lime"}`))
	}).Methods(http.MethodGet)

	r.PathPrefix("/line").Path("/webhook").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lineSignature := r.Header.Get("x-line-signature")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("failed io.ReadAll(r.Body)", "err", err)
			if errors.Is(err, application.ErrAuthInvalidArg) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := lineUsecase.VerifySignature(r.Context(), lineSignature, body); err != nil {
			slog.Error("failed lineUsecase.VerifySignature", "err", err)
			if errors.Is(err, application.ErrAuthInvalidArg) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.WriteString(w, "success")
	}).Methods(http.MethodPost)

	return httpadapter.New(r).ProxyWithContext, nil
}
