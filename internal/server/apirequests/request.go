// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apirequests

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

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

func (request *RequestImpl) GetURL() *url.URL {
	return request.request.URL
}

func (request *RequestImpl) GetVars() map[string]string {
	return mux.Vars(request.request)
}
