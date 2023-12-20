package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/gorilla/mux"
	"go.mod/entity"
	"go.mod/libs/rest"
	"go.mod/manager"
)

var (
	port         = ":6666"
	serverPorts  = []int{6667, 6668, 6669}
	currentIndex = 0
	mu           sync.Mutex
	requestCount int
)

func BalanceMiddleware(other http.HandlerFunc) http.HandlerFunc {
	fmt.Println("teste 01")
	mu.Lock()
	defer mu.Unlock()
	requestCount++

	fmt.Println("teste 02")
	if requestCount > 1 {
		return func(w http.ResponseWriter, r *http.Request) {
			size := rest.ReadRequest(w, r.Body)
			var ur *url.URL
			port := "6666"
			if size > 5 && requestCount > 1 {
				port = manager.GetNextPort()

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
			fmt.Println(v.Servers[0].URL)
			fmt.Println(requestCount)
			// Aqui estou resolvendo o bug.
			if port != "6666" {
				go http.ListenAndServe(":"+port, mux.NewRouter())
			}

			if ok := manager.IsServerHealthy(v.Servers[0].URL.String()); ok {
				proxy := httputil.NewSingleHostReverseProxy(v.Servers[0].URL)
				proxy.ServeHTTP(w, r)
			} else {
				if port == "6666" {

				}
				http.Error(w, "Servidor indisponível", http.StatusServiceUnavailable)
			}

		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		other.ServeHTTP(w, r)
	}
}
