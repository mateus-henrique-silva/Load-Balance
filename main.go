package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.mod/libs/rest"
	"go.mod/middleware"
)

var (
	port = []string{":6666", ":6667", ":6668", ":6669"}
)

var wg sync.WaitGroup

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		rest.Send(w, nil)
	})
	r.HandleFunc("/", middleware.BalanceMiddleware(balance)).Methods("GET")
	fmt.Println("O servidor esta escutando a porta...")
	fmt.Println(port)
	// Adicionando 1 à WaitGroup antes de iniciar o servidor
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(port[0], r)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(port[1], r); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(port[2], r); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(port[3], r); err != nil {
			fmt.Println(err)
		}
	}()
	// Aguardando a conclusão da goroutine do servidor
	wg.Wait()

}
