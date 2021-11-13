package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Name string
	ID   int64
}

type RecommendationsData struct {
	ProductList []Product
}

func Recommendations(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.
		New("recommendations.html").
		ParseFiles("resources/templates/recommendations.html")
	if err != nil {
		log.Fatal("Recommendations: "+err.Error(), err)
	}
	products := []Product{{Name: "kakashka", ID: 1}, {Name: "bukashka", ID: 2}, {Name: "kandibober", ID: 3}}

	data := RecommendationsData{
		ProductList: products,
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("SayHello: ", err)
	}
}
