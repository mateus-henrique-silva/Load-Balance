package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.mod/middleware"
)

const (
	port = ":6666"
)

var wg sync.WaitGroup

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.BalanceMiddleware(balance)).Methods("GET")
	fmt.Println("O servidor esta escutando a porta...")
	fmt.Println(port)
	// Adicionando 1 à WaitGroup antes de iniciar o servidor
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(port, r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Aguardando a conclusão da goroutine do servidor
	wg.Wait()

}
