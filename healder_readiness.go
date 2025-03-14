package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, struct{}{})

}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "something went wrong")
}
