package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewAlumno struct {
	Name    string
	IsEdit  bool
	Data    models.Alumno
	Widgets []models.Alumno
	UserId  string
}

var tmpla = template.Must(template.New("foo").Funcs(cfig.FuncMap).
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl",
		"web/alumno/index.html", "web/alumno/form.html"))

func AlumnoList(w http.ResponseWriter, req *http.Request) {

	session, _ := store.Get(req, "cookie-name")
	/*
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}*/

	/*alumno := models.Alumno{}

	alumnos, _ := alumno.FindAll(cfig.DB)

	for _, lis := range alumnos {
		fmt.Println(lis.ToString())
		fmt.Println("Matriculas: ", len(lis.Matriculas))
		if len(lis.Matriculas) > 0 {
			for _, d := range lis.Matriculas {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}*/

	// Create
	//cfig.DB.Create(&models.Alumno{Name: "Juan", City: "Juliaca"})
	lis := []models.Alumno{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewAlumno{
		Name:    "Alumno",
		Widgets: lis,
		UserId:  session.Values["user_id"].(string),
	}

	err := tmpla.ExecuteTemplate(w, "alumno/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Alumno
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
		d.Nombres = r.FormValue("nombres")
		d.Codigo = r.FormValue("codigo")
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
		http.Redirect(w, r, "/alumno/index", 301)
	}

	data := ViewAlumno{
		Name:   "Alumno",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpla.ExecuteTemplate(w, "alumno/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Alumno
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/alumno/index", 301)
}
