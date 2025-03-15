package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/piermou/rss_proj/internal/auth"
	"github.com/piermou/rss_proj/internal/database"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, struct{}{})

}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "something went wrong")
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responWithError(w, 400, fmt.Sprintf("error parsing JSON:", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("NOP create users: %v", err))
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, 200, databaseUserToUser(user))
}
