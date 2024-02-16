// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiResponses

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	writer http.ResponseWriter
}

func NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	return &JSONResponse{writer: w}
}

func (sender *JSONResponse) Send(statusCode int, data interface{}) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(sender.writer, "Error creating JSON writer", http.StatusInternalServerError)
		return
	}

	jsonData = append(jsonData, '\n')
	sender.writer.Header().Set("Content-Type", "application/json")
	sender.writer.WriteHeader(statusCode)
	_, err = sender.writer.Write(jsonData)

	if err != nil {
		http.Error(sender.writer, "Error writing JSON writer", http.StatusInternalServerError)
		return
	}

}

func (sender *JSONResponse) SendError(statusCode int, error string) {
	data := map[string]string{
		"error": error,
	}
	sender.Send(statusCode, data)
}

func (sender *JSONResponse) SendMethodNotSupportedError() {
	sender.SendError(http.StatusMethodNotAllowed, "Method is not supported.")
}
