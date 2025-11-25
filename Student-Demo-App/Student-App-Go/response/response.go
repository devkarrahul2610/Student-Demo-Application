package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JsonResponse(w http.ResponseWriter, status int, success bool, data interface{}, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res := APIResponse{
		Success: success,
		Data:    data,
		Error:   errMsg,
	}
	json.NewEncoder(w).Encode(res)
}

func ErrValidation(msg string) error {
	return errors.New(msg)
}
