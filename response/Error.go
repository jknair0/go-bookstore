package response

import "encoding/json"

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func CreateError(errorCode int, message string) *Error {
	return &Error{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func (e *Error) Encode() []byte {
	jsonStr, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return jsonStr
}

var (
	ErrUnknown              = CreateError(1, "Something went wrong")
	ErrInvalidRoute         = CreateError(2, "Invalid Path")
	ErrServerError          = CreateError(3, "Internal Server Error")
	ErrInvalidRequestParams = CreateError(4, "Invalid Request Params")
	ErrItemNotFound         = CreateError(5, "Item not found")
)
