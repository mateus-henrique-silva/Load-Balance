package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"

	"go.mod/entity"
	"go.mod/libs/rest"
)

var requestCount int
var mu sync.Mutex

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	mu.Lock()
	requestCount++
	mu.Unlock()
	return func(w http.ResponseWriter, r *http.Request) {
		size := rest.ReadRequest(w, r.Body)
		var ur *url.URL
		var port string
		if size > 10000 {
			for i := 0; i <= requestCount; i++ {
				n := 6666 + 1
				strconv.Itoa(n)
			}
		}

		ur, _ = url.Parse("http://localhost:" + port)
		v := entity.LoadBalancer{
			Servers: []*entity.BackendServer{
				{
					URL: ur,
				},
			},
		}
		proxy := httputil.NewSingleHostReverseProxy(v.Servers[0].URL)
		proxy.ServeHTTP(w, r)
		fmt.Println(size)
	}
}
