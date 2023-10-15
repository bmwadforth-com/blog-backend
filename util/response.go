package util

import (
	"errors"
)

type ApiResponse struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
	Error   string      `json:"error"`
}

func NewResponse(statusCode int, message string, body interface{}, err error) (int, ApiResponse) {
	if err != nil {
		if IsProduction {
			err = errors.New("an error has occurred")
		}

		return statusCode, ApiResponse{
			Message: message,
			Body:    body,
			Error:   err.Error(),
		}
	} else {
		return statusCode, ApiResponse{
			Message: message,
			Body:    body,
		}
	}
}

type DataResponse[T any] struct {
	Message string
	error   error
	Data    T
}

func NewDataResponse[T any](message string, data T) DataResponse[T] {
	return DataResponse[T]{
		Message: message,
		Data:    data,
	}
}

func (de *DataResponse[T]) SetData(data T) {
	de.Data = data
}

func (de *DataResponse[T]) SetError(error error) {
	SLogger.Errorf("an error has occurred: %v", error.Error())
	de.error = error
}

func (de *DataResponse[T]) GetError() error {
	return de.error
}

func (de *DataResponse[T]) GetErrorMessage() string {
	return de.error.Error()
}
