package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/libs/rest"
)

const (
	port = ":6666"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", balance).Methods("GET")
	fmt.Println("O servidor esta escutando a porta...")
	fmt.Println(port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		rest.ReturnStringError(err)
		return
	}

}
