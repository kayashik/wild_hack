package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func CommentsGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.
		New("comments.html").
		ParseFiles("resources/templates/comments.html")
	if err != nil {
		log.Fatal("Comments Get: ", err)
	}

	data := ViewData{
		TeamName: "@akshakak",
	}

	tplErr := tmpl.Execute(w, data)
	if tplErr != nil {
		log.Fatal("CommentsGet: ", err)
	}
}

func CommentsCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.Form.Get("comment")
	fmt.Fprintf(w, fmt.Sprintf("New comment: %s", comment))
}
