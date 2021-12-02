package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"wild_hack/controllers"
	"wild_hack/models"
)

func main() {
	e := echo.New()
	if e == nil {
		fmt.Sprintf("Съешь какашку")
	}

	// db connection
	err := models.Init()
	if err != nil {
		log.Fatal(err)
	}

	// routes
	http.HandleFunc("/", controllers.SayHello)                        // Устанавливаем роутер
	http.HandleFunc("/recommendations/", controllers.Recommendations) // Устанавливаем роутер
	http.HandleFunc("/comments/", controllers.CommentsGet)
	http.HandleFunc("/comments/create/", controllers.CommentsCreate)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./resources/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./resources/img"))))
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./resources/js"))))

	err = http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера http (пока)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
