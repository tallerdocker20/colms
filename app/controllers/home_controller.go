package controllers

import (
	"net/http"

	"github.com/202lp1/colms/models"
)

func Home(w http.ResponseWriter, req *http.Request) {

	d := models.Item{Title: "Ping", Notes: "Pong"}

	err := tmpl.ExecuteTemplate(w, "home/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
