package response

import "net/http"

func CreateSuccessHolder(data interface{}) *Holder {
	return &Holder{
		Data:  data,
		Error: nil,
	}
}

func CreateErrorHolder(error *Error) *Holder {
	return &Holder{
		Data:  nil,
		Error: error,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	_, _ = w.Write(CreateSuccessHolder(data).EncodeJson())
}

func WriteErrorResponse(w http.ResponseWriter, errMsg *Error) {
	_, _ = w.Write(CreateErrorHolder(errMsg).EncodeJson())
}
