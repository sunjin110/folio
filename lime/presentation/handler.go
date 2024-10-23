package presentation

import (
	"net/http"

	"github.com/sunjin110/folio/lime/application"
)

type HttpHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
	Hello(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	lineUsecase application.LineUsecase
}

func NewHttpHandler(lineUsecase application.LineUsecase) HttpHandler {
	return &httpHandler{
		lineUsecase: lineUsecase,
	}
}

func (h *httpHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "home"}`))
}

func (h *httpHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "lime"}`))
}
