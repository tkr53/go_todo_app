package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func ResponseJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	bobybytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("write error response error: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bobybytes); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
