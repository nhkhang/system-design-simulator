package pkg

import (
	"system-design-simulator/app/constants"
	"system-design-simulator/app/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constants.ResponseStatus, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Status:  int(responseStatus),
		Message: message,
		Data:    data,
	}
}
