package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("servidor rodando na porta 3000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Seja bem vindo ao meu servidor HTTP com Go")
	})

	http.ListenAndServe(":3000", mux)
}
