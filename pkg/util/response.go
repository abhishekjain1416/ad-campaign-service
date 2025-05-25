package util

import "github.com/abhishekjain1416/ad-campaign-service/pkg/dto"

func ParentResponse(messageCode string, message string) dto.ParentResponse {
	return dto.ParentResponse{
		MessageCode: messageCode,
		Message:     message,
	}
}

func SuccessResponse[T any](messageCode string, message string, data T) dto.SuccessResponse[T] {
	return dto.SuccessResponse[T]{
		ParentResponse: dto.ParentResponse{
			MessageCode: messageCode,
			Message:     message,
		},
		Data: data,
	}
}

func ErrorResponse(messageCode string, message string, errors []dto.Error) dto.ErrorResponse {
	return dto.ErrorResponse{
		ParentResponse: dto.ParentResponse{
			MessageCode: messageCode,
			Message:     message,
		},
		Error: errors,
	}
}
