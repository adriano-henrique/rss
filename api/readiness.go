package api

import "net/http"

type readinessResponse struct {
	Status string `json:"status"`
}

func ReadinessHandler(w http.ResponseWriter, req *http.Request) {
	respondWithJSON(w, http.StatusOK, readinessResponse{
		Status: "ok",
	})
}
