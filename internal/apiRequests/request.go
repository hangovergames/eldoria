// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiRequests

import (
	"net/http"
)

type Request interface {
	IsMethodGet() bool
}

type RequestImpl struct {
	request *http.Request
}

func NewRequest(
	request *http.Request,
) *RequestImpl {
	return &RequestImpl{request}
}

func (request *RequestImpl) IsMethodGet() bool {
	return request.request.Method == http.MethodGet
}
