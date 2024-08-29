package main

import (
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

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
