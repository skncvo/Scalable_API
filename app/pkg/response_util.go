package pkg

import (
	"github.com/skncvo/Scalable_API/app/constant"
	"github.com/skncvo/Scalable_API/app/domain/dto"
)

// interface{} : 동적 타입, 모든 타입을 담을 수 있는 빈 인터페이스
func NULL() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
