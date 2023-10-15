package util

import (
	"errors"
)

type ApiResponse[T any] struct {
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Error      string `json:"error"`
	statusCode int
}

func NewResponse[T any](statusCode int, message string, body T, err error) ApiResponse[T] {
	if err != nil {
		if IsProduction {
			err = errors.New("an error has occurred")
		}

		return ApiResponse[T]{
			Message:    message,
			Data:       body,
			Error:      err.Error(),
			statusCode: statusCode,
		}
	} else {
		return ApiResponse[T]{
			Message:    message,
			Data:       body,
			statusCode: statusCode,
		}
	}
}

func (ae *ApiResponse[T]) SetData(data T) {
	ae.Data = data
}

func (ae *ApiResponse[T]) SetError(error error) {
	SLogger.Errorf("an error has occurred: %v", error.Error())
	ae.Error = error.Error()
}

func (ae *ApiResponse[T]) GetError() string {
	return ae.Error
}

func (ae *ApiResponse[T]) SetStatusCode(statusCode int) {
	ae.statusCode = statusCode
}

func (ae *ApiResponse[T]) GetStatusCode() int {
	return ae.statusCode
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
