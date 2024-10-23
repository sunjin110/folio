package presentation

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/sunjin110/folio/lime/application"
	"github.com/sunjin110/folio/lime/config"
	"github.com/sunjin110/folio/lime/presentation/conv"
)

type HttpHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
	Hello(w http.ResponseWriter, r *http.Request)
	LineWebhook(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	lineChannelSecret string
	lineMessageClient *messaging_api.MessagingApiAPI
	lineUsecase       application.LineUsecase
}

func NewHttpHandler() (HttpHandler, error) {
	envConfig, err := config.NewEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("failed config.NewEnvConfig. err: %w", err)
	}

	lineMessageClient, err := messaging_api.NewMessagingApiAPI(envConfig.Line.ChannelToken)
	if err != nil {
		return nil, fmt.Errorf("failed messaging_api.NewMessagingApiAPI. err: %w", err)
	}

	lineUsecase := application.NewLineUsecase()

	return &httpHandler{
		lineChannelSecret: envConfig.Line.ChannelSecret,
		lineUsecase:       lineUsecase,
		lineMessageClient: lineMessageClient,
	}, nil
}

func (h *httpHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "home"}`))
}

func (h *httpHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "lime"}`))
}

func (h *httpHandler) LineWebhook(w http.ResponseWriter, r *http.Request) {
	req, err := webhook.ParseRequest(h.lineChannelSecret, r)
	if err != nil {
		if errors.Is(err, webhook.ErrInvalidSignature) {
			slog.Error("invalid line signature", "err", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		slog.Error("failed webhook.ParseRequest", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events := conv.ToLineEventsModel(req.Events)
	if err := h.lineUsecase.SaveContents(r.Context(), events); err != nil {
		slog.Error("failed lineUsecase.SaveContents", "err", err)
		return
	}
	io.WriteString(w, "success")
}
