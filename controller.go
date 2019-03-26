package goweb

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Data interface{}
}

func (c *Controller) ToJson(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	content, err := json.Marshal(c.Data)
	err_msg := "{err_code:1}"
	if err != nil {
		w.Write([]byte(err_msg))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}
