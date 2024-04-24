package main

import (
	"context"
	"log/slog"

	"github.com/sunjin110/folio/golio/presentation/http"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
)

func main() {
	slog.Info("hello world")

	cfg, err := httpconf.NewConfig()
	if err != nil {
		slog.Error("failed get config: %w", err)
	}

	slog.Info("config is ", "cfg", cfg)

	http.Serve(context.Background(), cfg)
}
