// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiresponses

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONResponse_Send(t *testing.T) {
	w := httptest.NewRecorder()
	sender := NewJSONResponse(w)

	data := map[string]string{"test": "value"}

	sender.Send(http.StatusOK, data)

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "{\"test\":\"value\"}\n"
	if body := w.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}

}

func TestJSONResponse_SendError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := NewJSONResponse(w)

	errorMessage := "An error occurred"
	statusCode := http.StatusBadRequest // You can choose any error status code

	sender.SendError(statusCode, errorMessage)

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	if status := w.Code; status != statusCode {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, statusCode)
	}

	// Check the response body for the correct error message
	expected := fmt.Sprintf("{\"error\":\"%s\"}\n", errorMessage)
	if body := w.Body.String(); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestJSONResponse_SendMethodNotSupportedError(t *testing.T) {
	w := httptest.NewRecorder()
	sender := NewJSONResponse(w)

	sender.SendMethodNotSupportedError()

	// Check the response content type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content type is wrong, got %s, want %s", contentType, "application/json")
	}

	// Check the status code
	expectedStatusCode := http.StatusMethodNotAllowed
	if status := w.Code; status != expectedStatusCode {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, expectedStatusCode)
	}

	// Check the response body for the correct error message
	expectedErrorMessage := "Method is not supported."
	expectedBody := fmt.Sprintf("{\"error\":\"%s\"}\n", expectedErrorMessage)
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}
