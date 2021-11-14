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

type Recommendation struct {
	Name 			string
	Result 			string
	Recommendation 	string
}

type RecommendationsData struct {
	ProductList 		[]Product
	RecommendationsList []Recommendation
}

func Recommendations(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.
		New("recommendations.html").
		ParseFiles("resources/templates/recommendations.html")
	if err != nil {
		log.Fatal("Recommendations: "+err.Error(), err)
	}
	products := []Product{{Name: "product 1", ID: 1}, {Name: "product 2", ID: 2}, {Name: "product 3", ID: 3}}
	recommendations := []Recommendation{
		{Name: "конверсия", Result: "+20%", Recommendation: "Вы в топ 15 продавцов на сайте"},
		{Name: "цена", Result: "топ 15% самых дорогих товаров", Recommendation: "Пополните склад"},
	}

	data := RecommendationsData{
		ProductList: products,
		RecommendationsList: recommendations,
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("SayHello: ", err)
	}
}
