package result

import "net/http"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewResponce(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

func NewBadRequest(msg string) Response {
	return Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}
}

func NewInternalServerError(msg string) Response {
	return Response{
		Code: http.StatusInternalServerError,
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
