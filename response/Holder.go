package response

import (
	"encoding/json"
)

type Holder struct {
	Data  interface{} `json:"data"`
	Error *Error      `json:"error"`
}

func (r *Holder) EncodeJson() []byte {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		errorResponseStr, _ := json.Marshal(CreateErrorHolder(ErrUnknown))
		return errorResponseStr
	}
	return jsonStr
}
