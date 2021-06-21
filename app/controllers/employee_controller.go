package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewEmployee struct {
	Name    string
	IsEdit  bool
	Data    models.Empleado
	Widgets []models.Empleado
}

var tmple = template.Must(template.New("foo").Funcs(cfig.FuncMap).
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl",
		"web/employee/index.html", "web/employee/form.html"))

func EmployeeList(w http.ResponseWriter, req *http.Request) {
	// Create
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})
	lis := []models.Empleado{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewEmployee{
		Name:    "Empleado",
		Widgets: lis,
	}

	err := tmple.ExecuteTemplate(w, "employee/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EmployeeForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Empleado
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Name = r.FormValue("name")
		d.City = r.FormValue("city")
		if id != "" {
			if err := cfig.DB.Save(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}

		} else {
			if err := cfig.DB.Create(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}
		}
		http.Redirect(w, r, "/employee/index", 301)
	}

	data := ViewEmployee{
		Name:   "Empleado",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmple.ExecuteTemplate(w, "employee/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EmployeeDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Empleado
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/employee/index", 301)
}
