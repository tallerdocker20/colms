package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewUser struct {
	Name    string
	IsEdit  bool
	Data    models.User
	Widgets []models.User
	UserId  string
}

var tmplu = template.Must(template.New("foo").Funcs(cfig.FuncMap).
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl",
		"web/user/index.html", "web/user/form.html", "web/user/login.html"))

func UserList(w http.ResponseWriter, req *http.Request) {

	session, _ := store.Get(req, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		//return
	}

	// Create
	//cfig.DB.Create(&models.User{Name: "Juan", City: "Juliaca"})
	lis := []models.User{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewUser{
		Name:    "User",
		Widgets: lis,
		UserId:  "", // session.Values["user_id"].(string),
	}

	err := tmplu.ExecuteTemplate(w, "user/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	user := models.User{}
	//var d models.User
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&user, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		user.PasswordConfirm = r.FormValue("password_confirm")
		if id != "" {
			/*if err := cfig.DB.Save(&user).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}*/

			err := user.UpdatePassword(cfig.DB)
			if err != nil {
				fmt.Println("Error in user.UpdatePassword()")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			/*if err := cfig.DB.Create(&user).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}*/
			err := user.Register(cfig.DB)
			if err != nil {
				fmt.Println("Error in user.Register()")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/user/index", 301)
	}

	data := ViewUser{
		Name:   "User",
		Data:   user,
		IsEdit: IsEdit,
	}

	err := tmplu.ExecuteTemplate(w, "user/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.User
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/user/index", 301)
}

func UserLoginForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)

	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	user := models.User{}
	//var d models.User
	if r.Method == "POST" {
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		err := user.IsAuthenticated(cfig.DB)
		if err != nil {
			fmt.Println("Error in user.Register()")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set user as authenticated
		session.Values["user_id"] = user.Id
		session.Values["authenticated"] = true
		session.Save(r, w)

		http.Redirect(w, r, "/user/index", 301)
	}

	data := ViewUser{
		Name: "User",
		Data: user,
	}

	err := tmplu.ExecuteTemplate(w, "user/loginformPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Values["user_id"] = ""
	session.Save(r, w)
	fmt.Fprintln(w, "thank you! see you")
}
