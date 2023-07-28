package api

import "net/http"

func ErrorHandler(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
