package main

import (
	"fmt"
	"net/http"

	"go.mod/entity"
	"go.mod/libs/rest"
)

func balance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ola mundo")
	rest.Send(w, entity.Main{
		Nome: "mateus henrique",
	})
	return
}
