package goweb

import (
	"net/http"
	"sync"
)

type Handler struct {
	h http.HandlerFunc
	m []MiddlewareType
}

type Router struct {
	m    map[string]*Handler
	pool sync.Pool
}

func NewRouter() *Router {
	m := make(map[string]*Handler)
	return &Router{
		m: m,
	}
}
func (r *Router) add(method string, path string, c http.HandlerFunc, mw ...MiddlewareType) {
	handler := &Handler{
		h: c,
	}
	middleWare := []MiddlewareType{}
	if len(mw) > 0 {
		for _, m := range mw {
			middleWare = append(middleWare, m)
		}
		handler.m = middleWare
	}
	r.m[method+"_"+path] = handler
}

func (r *Router) GET(path string, c http.HandlerFunc, mw ...MiddlewareType) {
	r.add("GET", path, c, mw...)
}
func (r *Router) POST(path string, c http.HandlerFunc, mw ...MiddlewareType) {
	r.add("POST", path, c, mw...)
}
func (r *Router) PUT(path string, c http.HandlerFunc, mw ...MiddlewareType) {
	r.add("PUT", path, c, mw...)
}
func (r *Router) OPTIONS(path string, c http.HandlerFunc, mw ...MiddlewareType) {
	r.add("OPTIONS", path, c, mw...)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	if h, ok := r.m[req.Method+"_"+url]; ok {
		handler := h.h
		if len(h.m) > 0 {
			for _, middle := range h.m {
				handler = middle(handler)
			}
		}
		handler(w, req)
	}
}
