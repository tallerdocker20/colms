package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewMatricula struct {
	Name    string
	IsEdit  bool
	Data    models.Matricula
	Widgets []models.Matricula
	Alumnos []models.Alumno
}

var tmplm = template.Must(template.New("foo").Funcs(cfig.FuncMap).
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl",
		"web/matricula/index.html", "web/matricula/form.html"))

func MatriculaList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Matricula{}
	if err := cfig.DB.Preload("Alumno").Find(&lis).Error; err != nil { // Preload("Alumno") carga los objetos Alumno relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewMatricula{
		Name:    "Matricula",
		Widgets: lis,
	}
	err := tmplm.ExecuteTemplate(w, "matricula/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MatriculaForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Matricula
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	alumno := models.Alumno{}
	alumnos, _ := alumno.GetAll(cfig.DB) // para mostrar los alumnos en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Semestre = r.FormValue("semestre")
		//n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.AlumnoId = r.FormValue("alumno_id") //n
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
		http.Redirect(w, r, "/matricula/index", 301)
	}

	data := ViewMatricula{
		Name:    "Matricula",
		Data:    d,
		IsEdit:  IsEdit,
		Alumnos: alumnos,
	}

	err := tmplm.ExecuteTemplate(w, "matricula/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MatriculaDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Matricula
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/matricula/index", 301)
}
