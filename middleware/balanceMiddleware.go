package middleware

import (
	"fmt"
	"net/http"

	"go.mod/libs/rest"
)

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		size := rest.ReadRequest(w, r.Body)
		fmt.Println(size)
	}
}
