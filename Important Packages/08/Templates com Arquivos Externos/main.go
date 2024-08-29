package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("content.html").ParseFiles(templates...))
		tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})

		err := tmp.Execute(os.Stdout, Cursos{
			{"Golang", 40},
			{"Java", 30},
			{"Python", 25},
			{"Ruby", 20},
			{"JavaScript", 15},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8000", nil)
}
