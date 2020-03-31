package network

import (
	"encoding/json"
	"github.com/google/logger"
	"net/http"
)

func CreateAnswer(w http.ResponseWriter, statusCode int, value interface{})  {
	encoded , err := json.Marshal(value)
	if err != nil {
		logger.Errorf("Error marhalling json %v", err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(encoded)
	if err != nil {
		logger.Errorf("Error writing answer %v", err)
	}
}
