package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct{
	TeamName string
	Url string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.
		New("homePage.html").
		ParseFiles("resources/templates/homePage.html")
	if err != nil {
		log.Fatal("SayHello: ", err)
	}

	data := ViewData{
		TeamName: "@akshakak",
		Url: "/recommendations",
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("SayHello: ", err)
	}
}