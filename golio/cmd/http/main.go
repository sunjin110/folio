package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/sunjin110/folio/golio/presentation/http"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
)

func main() {
	cfg, err := httpconf.NewConfig()
	if err != nil {
		slog.Error("failed get config: %w", err)
		os.Exit(1)
	}

	slog.Info("config is ", "cfg", cfg)

	if err := http.Serve(context.Background(), cfg); err != nil {
		slog.Error("failed http.Serve: %w", err)
		os.Exit(1)
	}
}
