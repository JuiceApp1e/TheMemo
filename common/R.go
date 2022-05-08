package common

type R struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func Ok(data interface{}) (r R) {
	r.Code = 0
	r.Data = data
	return r
}

func Error(data interface{}, msg string) (r R) {
	r.Code = 1
	r.ErrorMsg = msg
	r.Data = data
	return r
}
