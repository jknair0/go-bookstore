package response

import "encoding/json"

type Error struct {
	ErrorCode ApiErrorCode `json:"error_code"`
	Message   string       `json:"message"`
}

func NewError(code ApiErrorCode, message string) *Error {
	return &Error{
		ErrorCode: code,
		Message:   message,
	}
}

func NewUnknownError(err error) *Error {
	return NewUnknownErrorFromMsg(err.Error())
}

func NewUnknownErrorFromMsg(message string) *Error {
	return &Error{
		ErrorCode: ErrUnknown,
		Message:   message,
	}
}

func NewErrorFromCode(errorCode ApiErrorCode) *Error {
	return &Error{
		ErrorCode: errorCode,
		Message:   ErrorMessage[errorCode],
	}
}

func (e *Error) Encode() []byte {
	jsonStr, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return jsonStr
}

type ApiErrorCode int

const (
	ErrUnknown ApiErrorCode = iota
	ErrInvalidRoute
	ErrServerError
	ErrInvalidRequestFormat
	ErrInvalidRequestBody
	ErrItemNotFound
)

var ErrorMessage = map[ApiErrorCode]string{
	ErrUnknown:              "Something went wrong",
	ErrInvalidRoute:         "Invalid Path",
	ErrServerError:          "Internal Server Error",
	ErrInvalidRequestFormat: "Invalid Request Format",
	ErrInvalidRequestBody:   "Invalid Request Body",
	ErrItemNotFound:         "Item not found",
}
