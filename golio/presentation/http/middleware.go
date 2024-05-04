package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sunjin110/folio/golio/usecase"
)

// authMW 認証する
func AuthMW(authUsecase usecase.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/auth/google-oauth") {
				// 認証の必要なし
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("access_token")
			if err != nil {
				if errors.Is(err, http.ErrNoCookie) {
					fmt.Println("見つからなかった")
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}
				slog.Error("failed get cookie", "err", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}

			userSession, err := authUsecase.GetSessionInfoFromToken(r.Context(), cookie.Value)
			if err != nil {
				slog.Error("failed get user session", "err", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
			if userSession == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
