package model

import "encoding/json"

type ResponseHolder struct {
	Data  interface{} `json:"data"`
	Error *string      `json:"error"`
}

func CreateSuccessResponseHolder(data interface{}) *ResponseHolder {
	return &ResponseHolder{
		Data:  data,
		Error: nil,
	}
}

func CreateErrorResponseHolder(error string) *ResponseHolder {
	return &ResponseHolder{
		Data:  nil,
		Error: &error,
	}
}

var UnknownErrorResponseHolder = CreateErrorResponseHolder(UNKNOWN_REASON_ERROR_MESSAGE)

func (r *ResponseHolder) EncodeJson() []byte {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		errorResponseStr, _ := json.Marshal(UnknownErrorResponseHolder)
		return errorResponseStr
	}
	return jsonStr
}
