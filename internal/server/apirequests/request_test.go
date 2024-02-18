// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apirequests

import (
	"net/http"
	"testing"
)

func TestRequestImpl_IsMethodGet(t *testing.T) {

	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{"GET Method", http.MethodGet, true},
		{"POST Method", http.MethodPost, false},
		{"PUT Method", http.MethodPut, false},
		{"DELETE Method", http.MethodDelete, false},
		{"PATCH Method", http.MethodPatch, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, "/", nil)
			if err != nil {
				t.Fatal("Creating request failed:", err)
			}

			requestImpl := NewRequest(req)
			if got := requestImpl.IsMethodGet(); got != test.expected {
				t.Errorf("IsMethodGet() = %v, want %v for method %s", got, test.expected, test.method)
			}
		})
	}

}
