package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Name string
	ID   int64
}

type Recommendation struct {
	Name 			string
	Result 			string
	Recommendation 	string
	Compare			string
}

type RecommendationsData struct {
	ProductList 		[]Product
	RecommendationsList []Recommendation
	ProductID 			int64
	PlotImageURL		string
}

func Recommendations(w http.ResponseWriter, r *http.Request) {
	var productID int64 = 1
	productIDs, ok := r.URL.Query()["id"]

	if ok && len(productIDs) > 0 {
		productID, _ = strconv.ParseInt(productIDs[0], 10, 64)
	}

	tmpl, err := template.
		New("recommendations.html").
		ParseFiles("resources/templates/recommendations.html")
	if err != nil {
		log.Fatal("Recommendations: "+err.Error(), err)
	}
	products := []Product{{Name: "product 1", ID: 1}, {Name: "product 2", ID: 2}, {Name: "product 3", ID: 3}}
	recommendations := []Recommendation{
		{Name: "конверсия", Result: "+20%", Compare: "Вы в топ 15 продавцов на сайте", Recommendation: "Продолжайте в том же духе"},
		{Name: "цена", Result: "топ 15% самых дорогих товаров", Compare: "Вы в топ 15 продавцов на сайте",  Recommendation: "Пополните склад"},
	}

	data := RecommendationsData{
		ProductList: products,
		RecommendationsList: recommendations,
		ProductID: productID,
		PlotImageURL: fmt.Sprintf("../img/plot-%d.png", productID),
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("Recommendations template err: " + err.Error(), err)
	}
}
