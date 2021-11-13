package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct{
	Title string
	Message string
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
		Title: "Wild Hack",
		Message: "Click here to get recommendations",
		Url: "http://localhost:8080/recommendations",
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("SayHello: ", err)
	}
}