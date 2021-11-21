package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Name 	   string
	ID   	   int64
	IsSelected bool
}

type Recommendation struct {
	Name 			string
	Result 			string
	Recommendation 	string
	Compare			string
}

type PlotRecom struct {
	Recom  string
}

type RecommendationsData struct {
	ProductList 		[]Product
	RecommendationsList []Recommendation
	ProductID 			int64
	PlotImageURL		string
	PlotRecommendation  string
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
	products := []Product{
		{Name: "02C451A73", ID: 1, IsSelected: 1 == productID}, // 02C451A73FDD41CDD05F1E33E4BC60FC
		{Name: "3AD6487EB", ID: 2, IsSelected: 2 == productID}, // 3AD6487EBF165BDE4EECE355B64A80A8
		{Name: "FC7C2D6A3", ID: 3, IsSelected: 3 == productID}, // FC7C2D6A3572340064B2B06935AC200C
	}
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

	plotResults := map [int64]PlotRecom{
		1: {Recom: "Продажи могут значительно вырасти, рекомендуем следить за товарными остатками"},
		2: {Recom: "Продажи значительно вырастут, рекомендуем пополнить товарными остатками"},
		3: {Recom: "Продажи товаров растут не линейно, продолжайте в том же духе"},
	}

	data := RecommendationsData{
		ProductList: products,
		RecommendationsList: recommendations[productID],
		ProductID: productID,
		PlotImageURL: fmt.Sprintf("../img/plot-%d.png", productID),
		PlotRecommendation: plotResults[productID].Recom,
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("Recommendations template err: " + err.Error(), err)
	}
}
