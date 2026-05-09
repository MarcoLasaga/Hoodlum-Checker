package main

import (
	"html/template"
	"log"
)

var (
	landingTmpl *template.Template
	quizTmpl    *template.Template
	resultTmpl  *template.Template
)

func loadTemplates() {
	var err error

	landingTmpl, err = template.ParseFiles("templates/landing.html")
	if err != nil {
		log.Fatal(err)
	}

	quizTmpl, err = template.ParseFiles("templates/quiz.html")
	if err != nil {
		log.Fatal(err)
	}

	resultTmpl, err = template.ParseFiles("templates/result.html")
	if err != nil {
		log.Fatal(err)
	}
}
