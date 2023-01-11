package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) broker(w http.ResponseWriter, _ *http.Request) {
	type JSONResponse struct {
		Message string `json:"message"`
		Error   bool   `json:"error"`
		Data    any    `json:"data,omitempty"`
	}

	response := JSONResponse{
		Message: "Hit the broker.",
		Error:   false,
	}

	b, _ := json.MarshalIndent(response, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_, _ = w.Write(b)
}
