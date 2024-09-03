package errors

import (
    // "log"
    "net/http"
)

type GenericError struct {
    Message string `json:"message"`
    StatusCode int `json:"status_code"`
}

func newError(message string, statusCode int) *GenericError {
    return &GenericError{
        Message: message,
        StatusCode: statusCode,
    }
}

func sendError(w http.ResponseWriter, err *GenericError) {
    http.Error(w, err.Message, err.StatusCode)
}

func InvalidRequestError(w http.ResponseWriter) {
    err := newError("Invalid request", http.StatusBadRequest)
    sendError(w, err)
}

func InternalServerError(w http.ResponseWriter, inErr error) {
    // log.Fatal(inErr)
    err := newError("Internal server error", http.StatusInternalServerError)
    sendError(w, err)
}

func BadRequestError(w http.ResponseWriter, message string) {
    err := newError(message, http.StatusBadRequest)
    sendError(w, err)
}

func UnauthorizedError(w http.ResponseWriter) {
    err := newError("Unauthorized", http.StatusUnauthorized)
    sendError(w, err)
}
