package main

import (
	"log"
	"net/http"

	"wild_hack/controllers"
)

func main() {
	http.HandleFunc("/", controllers.SayHello) // Устанавливаем роутер
	http.HandleFunc("/recommendations", controllers.Recommendations) // Устанавливаем роутер


	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера http (пока)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}