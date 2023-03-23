package result

import "net/http"

type Response struct {
	Code int
	Msg  string
}

func NewResponce(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

func Success(msg string) Response {
	return Response{
		Code: http.StatusOK,
		Msg:  msg,
	}
}

func (r *Response) IsErr() bool {
	return r.Code >= 400 || r.Code == 0
}
