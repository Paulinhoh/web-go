package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("servidor rodando na porta 3000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Seja bem vindo ao meu servidor HTTP com Go")

		t, err := template.ParseFiles("./index.html")
		if err != nil {
			http.Error(w, "aconteceu um erro ao executar essa página", http.StatusInternalServerError)
			return
		}
		t.ExecuteTemplate(w, "index", nil)
	})

	http.ListenAndServe(":3000", mux)
}
