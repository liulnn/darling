package darling

import (
	"net/http"
)

type Response struct {
	Out         http.ResponseWriter
	StatusCode  int
	Header      map[string]string
	ContentType string
	Content     []byte
}

func (resp *Response) Write() {
	resp.Out.WriteHeader(resp.StatusCode)
	resp.Out.Header().Add("Content-Type", resp.ContentType)
	for k, v := range resp.Header {
		resp.Out.Header().Add(k, v)
	}
	resp.Out.Write(resp.Content)
}
