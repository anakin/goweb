package goweb

import (
	"net/http"
	"time"
)

type Response struct {
	http.ResponseWriter
	Started bool
	Status  int
	Elapsed time.Duration
}

type Context struct {
	Request        *http.Request
	ResponseWriter *Response
	_xsrfToken     string
}
