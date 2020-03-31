package network

import (
	"github.com/google/logger"
	"net/http"
)

func CreateAnswer(w http.ResponseWriter, statusCode int, data []byte)  {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	_, err := w.Write(data)
	if err != nil {
		logger.Errorf("Error writing answer %v", err)
	}
}
