package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles("./cmd/exp/layout1.html", "./cmd/exp/layout2.html", "./cmd/exp/home.html", "./cmd/exp/footer.html", "./cmd/exp/header.html")
	if err != nil {
		panic(err)
	}

	// data := TemplateData{Nome: "paulinho", Idade: 24}
	err = t.ExecuteTemplate(os.Stdout, "layout1.html", "2026")
	if err != nil {
		panic(err)
	}
}
