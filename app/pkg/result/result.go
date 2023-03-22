package result

import "net/http"

type Responce struct {
	Code int
	Msg  string
}

func NewResponce(code int, msg string) Responce {
	return Responce{
		Code: code,
		Msg:  msg,
	}
}

func Success(msg string) Responce {
	return Responce{
		Code: http.StatusOK,
		Msg:  msg,
	}
}

func (r *Responce) IsErr() bool {
	return r.Code >= 400 || r.Code == 0
}
