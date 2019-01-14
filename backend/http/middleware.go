package http

import (
	"encoding/json"
	"net/http"

	"github.com/luanngominh/mnotes/backend/util"
)

type unauthorize struct {
	Err string `json:"error"`
}

//NoteMiddleware check user login
func NoteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if err := util.VerifyToken(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			js, _ := json.Marshal(unauthorize{Err: err.Error()})
			w.Write(js)
			return
		}
		next.ServeHTTP(w, req)
	})
}
