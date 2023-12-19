package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"go.mod/entity"
	"go.mod/libs/rest"
)

var requestCount int
var mu sync.Mutex

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	fmt.Println("teste 01")
	mu.Lock()
	requestCount++
	mu.Unlock()
	fmt.Println("teste 02")
	return func(w http.ResponseWriter, r *http.Request) {
		size := rest.ReadRequest(w, r.Body)
		var ur *url.URL
		var port string
		if size > 5 {
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
		fmt.Println(size)
		fmt.Println(v)
		fmt.Println(port)
		// Aqui estou resolvendo o bug.
		// proxy := httputil.NewSingleHostReverseProxy(v.Servers[0].URL)
		// proxy.ServeHTTP(w, r)

	}
}
