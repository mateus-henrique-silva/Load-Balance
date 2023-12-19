package middleware

import (
	"fmt"
	"net/http"
	"net/url"

	"go.mod/entity"
	"go.mod/libs/rest"
)

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		size := rest.ReadRequest(w, r.Body)
		ur, _ := url.Parse("http://localhost:6667")
		v := entity.LoadBalancer{
			Servers: []*entity.BackendServer{
				&entity.BackendServer{
					URL: ur,
				},
			},
		}

		fmt.Println(size)
	}
}
