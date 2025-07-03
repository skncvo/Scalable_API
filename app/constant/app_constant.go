package constant

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	// iota 사용 -> success =1 , DataNotFound =2, ...
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

// 메소드
func (r ResponseStatus) GetResponseStatus() string {
	//return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
	statuses := [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}
	if r >= 1 && int(r) <= len(statuses) {
		return statuses[r-1]
	}
	return "UNKNOWN"
}

func (r ResponseStatus) GetResponseMessage() string {
	statuses := [...]string{"Success", "Data_NotFound", "Unknown_Error", "Invalid_Request", "Unauthorized"}
	if r >= 1 && int(r) <= len(statuses) {
		return statuses[r-1]
	}
	return "Unknown"
}
