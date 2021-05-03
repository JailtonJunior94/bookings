package responses

import (
	"fmt"
	"net/http"
)

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func newHttpResponse(statusCode int, data interface{}) *HttpResponse {
	return &HttpResponse{StatusCode: statusCode, Data: data}
}

func formatError(message interface{}) map[string]string {
	mapError := make(map[string]string)
	mapError["error"] = fmt.Sprintf("%v", message)

	return mapError
}

func Ok(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusOK, data)
}

func Created(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusCreated, data)
}

func NoContent() *HttpResponse {
	return newHttpResponse(http.StatusNoContent, nil)
}

func BadRequest(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusBadRequest, formatError(data))
}

func Unauthorized(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusUnauthorized, formatError("Token inválido ou expirado"))
}

func NotFound(data interface{}) *HttpResponse {
	return newHttpResponse(http.StatusNotFound, formatError("Recurso não encontrado"))
}

func ServerError() *HttpResponse {
	return newHttpResponse(http.StatusInternalServerError, formatError("Ocorreu um erro inesperado"))
}
