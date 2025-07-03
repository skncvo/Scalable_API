package dto

// 제네릭 T 사용 : T 는 어떤 타입이든 될 수 있음
type ApiResponse[T any] struct {
	// 상태 코드 : success, invalid ...
	ResponseKey string `json:"response_key"`
	// 사용자 친화적 메시지
	ResponseMessage string `json:"response_message"`
	// 응답데이터
	Data T `json:"data"`
}
