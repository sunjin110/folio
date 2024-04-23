package main

import (
	"context"
	"log/slog"

	"github.com/sunjin110/folio/golio/presentation/http"
)

func main() {
	slog.Info("hello world")

	http.Serve(context.Background())
}
