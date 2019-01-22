package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/luanngominh/mnotes/backend/util"
)

type unauthorize struct {
	Err string `json:"error"`
}

// type AuthKey *jwt.

//NoteMiddleware check user login
func NoteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		authInfo, err := util.VerifyToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			js, _ := json.Marshal(unauthorize{Err: err.Error()})
			w.Write(js)
			return
		}

		ctx := req.Context()
		ctx = context.WithValue(ctx, "auth_info", authInfo)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
