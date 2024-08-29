package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, []Curso{
		{"Golang", 40},
		{"Java", 30},
		{"Python", 25},
		{"Ruby", 20},
		{"JavaScript", 15},
	})

	if err != nil {
		panic(err)
	}
}
