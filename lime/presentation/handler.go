package presentation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	aws_cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/sunjin110/folio/lime/application"
	"github.com/sunjin110/folio/lime/config"
	"github.com/sunjin110/folio/lime/infrastructure/repository"
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

func NewHttpHandler(ctx context.Context) (HttpHandler, error) {
	envConfig, err := config.NewEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("failed config.NewEnvConfig. err: %w", err)
	}

	lineMessageClient, err := messaging_api.NewMessagingApiAPI(envConfig.Line.ChannelToken)
	if err != nil {
		return nil, fmt.Errorf("failed messaging_api.NewMessagingApiAPI. err: %w", err)
	}

	lineClient, err := linebot.New(envConfig.Line.ChannelSecret, envConfig.Line.ChannelToken)
	if err != nil {
		return nil, fmt.Errorf("failed linebot.New. err: %w", err)
	}

	awsCfg, err := aws_cfg.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed aws_cfg.LoadDefaultConfig. err: %w", err)
	}

	lineContentRepo := repository.NewLineContent(lineClient)
	lineMessageRepo := repository.NewLineMessage(lineMessageClient)
	storageRepo := repository.NewStorage(s3.NewFromConfig(awsCfg), envConfig.MediaS3BucketName)

	lineUsecase := application.NewLineUsecase(lineContentRepo, lineMessageRepo, storageRepo)

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
