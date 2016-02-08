package response

import (
	"encoding/json"
	"net/http"
)

var (
	BadRequest   = Empty(http.StatusBadRequest)
	Unauthorized = Empty(http.StatusUnauthorized)
	NotFound     = Empty(http.StatusNotFound)
	ServerError  = Empty(http.StatusInternalServerError)
)

type Response interface {
	WriteTo(out http.ResponseWriter)
}

type NormalResponse struct {
	status int
	body   []byte
	header http.Header
}

func (r *NormalResponse) WriteTo(out http.ResponseWriter) {
	header := out.Header()
	for k, v := range r.header {
		header[k] = v
	}
	out.WriteHeader(r.status)
	out.Write(r.body)
}

func (r *NormalResponse) Header(key, value string) *NormalResponse {
	r.header.Set(key, value)
	return r
}

func Empty(status int) *NormalResponse {
	return Respond(status, nil)
}

func Json(status int, body interface{}) *NormalResponse {
	return Respond(status, body).Header("Content-Type", "application/json")
}

func Error(message string, err error) *NormalResponse {
	return ServerError
}

func Respond(status int, body interface{}) *NormalResponse {
	var b []byte
	var err error
	switch t := body.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	default:
		if body != nil {
			if b, err = json.Marshal(body); err != nil {
				return &NormalResponse{
					status: http.StatusInternalServerError,
				}
			}
		}
	}
	return &NormalResponse{
		body:   b,
		status: status,
		header: make(http.Header),
	}
}
