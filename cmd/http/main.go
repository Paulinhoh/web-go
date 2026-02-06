package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	w.Header()["Date"] = nil // suprimir esse cabeçalho

	fmt.Fprint(w, "listagem de notas")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "nota não encontrada", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "exibindo uma nota especifica "+id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// rejeitar a reqisição
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "criando uma nova nota")
}

func main() {
	fmt.Println("servidor rodando na porta 8080")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":8080", mux)
}
