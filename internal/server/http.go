// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package server

import "encoding/json"

const (
	InternalServerError        = "server_error"
	InternalServerErrorMessage = "Внутренняя ошибка сервера"

	Unauthorized        = "unauthorized"
	UnauthorizedMessage = "Вы не авторизованы"

	BadRequest        = "bad_request"
	BadRequestMessage = "Переданы некорректные данные"
)

type HTTPError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	RequestID    string `json:"request_id"`
}

type HTTPAuthSuccess struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewHTTPError(errorCode, errorMessage, requestID string) []byte {
	data, _ := json.Marshal(HTTPError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		RequestID:    requestID,
	})
	return data
}

func NewHTTPAuthSuccess(accessToken, refreshToken string) []byte {
	data, _ := json.Marshal(HTTPAuthSuccess{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	return data
}

func NewHTTPInternalServerError(requestID string) []byte {
	data, _ := json.Marshal(HTTPError{
		ErrorCode:    InternalServerError,
		ErrorMessage: InternalServerErrorMessage,
		RequestID:    requestID,
	})
	return data
}

func NewHTTPUnauthorizedError(requestID string) []byte {
	data, _ := json.Marshal(HTTPError{
		ErrorCode:    Unauthorized,
		ErrorMessage: UnauthorizedMessage,
		RequestID:    requestID,
	})
	return data
}

func NewHTTPBadRequestError(requestID string) []byte {
	data, _ := json.Marshal(HTTPError{
		ErrorCode:    BadRequest,
		ErrorMessage: BadRequestMessage,
		RequestID:    requestID,
	})
	return data
}
