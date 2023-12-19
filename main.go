package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/libs/rest"
	"go.mod/middleware"
)

const (
	port = ":6666"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.BalanceMiddleware(balance)).Methods("GET")
	fmt.Println("O servidor esta escutando a porta...")
	fmt.Println(port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		rest.ReturnStringError(err)
		return
	}

}
