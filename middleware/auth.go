package goweb

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//sess, _ := globalSessions.SessionStart(w, r)
		//cookie, err := r.Cookie("gosessid")
		//if err != nil {
		//	log.Println(err)
		//}
		//sessid := cookie.Value

		next.ServeHTTP(w, r)
	}
}
