package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

var tmpli = template.Must(template.New("foo").Funcs(cfig.FuncMap).
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl",
		"web/item/index.html"))

func ItemList(w http.ResponseWriter, req *http.Request) {

	//t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")

	//vars := mux.Vars(req)
	//fmt.Println("id=", vars["id"])
	//fmt.Fprintf(w, "tablaget page ", vars["id"])
	// you access the cached templates with the defined name, not the filename
	d := models.Item{Title: "Sean222", Notes: "nnn"}

	err := tmpli.ExecuteTemplate(w, "item/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ItemTemp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "item list page ooooo ")
}
