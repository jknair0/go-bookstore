package response

import (
	"encoding/json"
)

type Holder struct {
	Data  interface{} `json:"data"`
	Error *Error      `json:"error"`
}

func NewHolder(data interface{}, err *Error) *Holder {
	return &Holder{
		Data:  data,
		Error: err,
	}
}

func NewEncodedSuccessHolder(data interface{}) []byte {
	return NewHolder(data, nil).EncodeJson()
}

func NewEncodedErrorHolder(err *Error) []byte {
	return NewHolder(nil, err).EncodeJson()
}

func (r *Holder) EncodeJson() []byte {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return jsonStr
}
