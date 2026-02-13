package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	files := []string{
		"./views/templates/base.html",
		"./views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "nota não encontrada", http.StatusNotFound)
		return
	}

	files := []string{
		"./views/templates/base.html",
		"./views/templates/pages/note-view.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/templates/base.html",
		"./views/templates/pages/note-new.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
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

	staticHandler := http.FileServer(http.Dir("./views/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":8080", mux)
}
