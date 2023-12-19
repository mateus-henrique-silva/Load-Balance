package middleware

import (
	"net/http"

	"go.mod/libs/rest"
)

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.ReadRequest(w, r.Body)
	}
}
