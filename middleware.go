package goweb

import "net/http"

type MiddlewareType func(next http.HandlerFunc) http.HandlerFunc
