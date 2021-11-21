package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"wild_hack/controllers"
)

func main() {
	e := echo.New();
	if e == nil {
		fmt.Sprintf("Съешь какашку")
	}
	http.HandleFunc("/", controllers.SayHello)                            // Устанавливаем роутер
	http.HandleFunc("/recommendations/", controllers.Recommendations) // Устанавливаем роутер
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./resources/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./resources/img"))))
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./resources/js"))))

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера http (пока)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
