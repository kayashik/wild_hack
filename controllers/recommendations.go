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
	products := []Product{{Name: "02C451A73FDD41CDD05F1E33E4BC60FC", ID: 1}, {Name: "3AD6487EBF165BDE4EECE355B64A80A8", ID: 2}, {Name: "FC7C2D6A3572340064B2B06935AC200C", ID: 3}}
	recommendations := map [int64] []Recommendation{
		1: {
			{Name: "конверсия", Result: "+2%", Compare: "Конверсия ниже, чем у 85% аналогичных товаров у других продавцов", Recommendation: "Рекомендуем пересмотреть цену"},
			{Name: "цена", Result: "1744", Compare: "на 3% выше средней в сравнении а аналогичными товарами",  Recommendation: "Рекомендуем сделать скидку"},
		},
		2: {
			{Name: "конверсия", Result: "9,75%", Compare: "Продукт входит в топ-15% по уровню конверсии", Recommendation: "Продолжайте в том же духе"},
			{Name: "цена", Result: "466", Compare: "на 2% выше среднейв сравнении с аналогичными товарами",  Recommendation: "Для увеличения продаж рекомендуется проводить акции"},
		},
		3: {
			{Name: "конверсия", Result: "+7%", Compare: "Продукт входит в топ-15% по уровню конверсии", Recommendation: "Продолжайте в том же духе"},
			{Name: "цена", Result: "1933", Compare: "на 2,4% выше среднейв сравнении с аналогичными товарами",  Recommendation: "Рекомендуем сделать скидку"},
		},
	}

	data := RecommendationsData{
		ProductList: products,
		RecommendationsList: recommendations[productID],
		ProductID: productID,
		PlotImageURL: fmt.Sprintf("../img/plot-%d.png", productID),
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("Recommendations template err: " + err.Error(), err)
	}
}
