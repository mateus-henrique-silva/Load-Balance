package main

import (
	"fmt"
	"net/http"
)

func balance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ola mundo")
	return
}
