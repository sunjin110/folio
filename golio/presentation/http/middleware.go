package http

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/usecase"
)

// authMW 認証する
func AuthMW(authUsecase usecase.Auth, userRepo repository.User) mux.MiddlewareFunc {
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
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}
				slog.Error("failed get cookie", "err", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}

			userSession, err := authUsecase.GetSessionInfoFromToken(r.Context(), cookie.Value)
			if err != nil {
				if errors.Is(err, usecase.ErrNotFound) {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}
				slog.Error("failed get user session", "err", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}

			// check
			if userSession.AccessTokenExpireTime.Before(time.Now()) {
				// userInfoにアクセスしてrefresh_tokenを取得する必要がある

				user, err := userRepo.Get(r.Context(), userSession.Email)
				if err != nil {
					if errors.Is(err, repository.ErrNotFound) {
						slog.Info("no registed email", "email", userSession.Email)
						http.Error(w, "unauthorized", http.StatusUnauthorized)
						return
					}
					slog.Error("failed get user info", "err", err, "email", userSession.Email)
					http.Error(w, "internal error", http.StatusInternalServerError)
					return
				}

				refreshedUserSession, err := authUsecase.RefreshSession(r.Context(), user.RefreshToken, user.Email)
				if err != nil {
					slog.Error("failed refresh token", "err", err)
					http.Error(w, "refresh token error", http.StatusUnauthorized)
					return
				}

				// CookieのaccessTokenを更新
				http.SetCookie(w, &http.Cookie{
					Name:     "access_token",
					Value:    refreshedUserSession.AccessToken,
					HttpOnly: true,
					Secure:   true,
					Path:     "/",
					SameSite: http.SameSiteNoneMode,
				})
			}

			next.ServeHTTP(w, r)
		})
	}
}
