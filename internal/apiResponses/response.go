// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiResponses

type Response interface {
	Send(statusCode int, data interface{})
	SendError(statusCode int, error string)
	SendMethodNotSupportedError()
}
