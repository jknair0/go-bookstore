package response

import "net/http"

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	_, _ = w.Write(NewEncodedSuccessHolder(data))
}

func WriteUnknownErrorMsg(w http.ResponseWriter, message string) {
	_, _ = w.Write(NewEncodedErrorHolder(NewUnknownErrorFromMsg(message)))
}

func WriteUnknownError(w http.ResponseWriter, err error) {
	_, _ = w.Write(NewEncodedErrorHolder(NewUnknownError(err)))
}

func WriteErrorCodeResponse(w http.ResponseWriter, errorCode ApiErrorCode) {
	_, _ = w.Write(NewEncodedErrorHolder(NewErrorFromCode(errorCode)))
}

func WriteErrorCodeCustomMessageResponse(w http.ResponseWriter, errorCode ApiErrorCode, message string) {
	_, _ = w.Write(NewEncodedErrorHolder(NewError(errorCode, message)))
}
