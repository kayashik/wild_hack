package controllers

import (
	"fmt"
	"net/http"
)

func Recommendations(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "It's reccomendations page!")
}
