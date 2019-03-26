package goweb

import (
	"net/http"
)

type GoWeb struct {
}

func New() *GoWeb {
	return &GoWeb{}
}

func (g *GoWeb) Run(mux *Router) {
	http.ListenAndServe(":8081", mux)
}
