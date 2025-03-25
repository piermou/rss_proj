package main

import (
	"fmt"
	"net/http"

	"github.com/piermou/rss_proj/internal/auth"
	"github.com/piermou/rss_proj/internal/database"
)

type autheHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler autheHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUSerByAPIKey(r.Context(), apiKey)
		if err != nil {
			responWithError(w, 403, fmt.Sprintf("auth error : %v", err))
			return
		}
		handler(w, r, user)
	}

}
