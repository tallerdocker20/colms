package routes

import (
	"github.com/202lp1/colms/controllers"
	"github.com/gorilla/mux" //gin or chi u other
)

func RoutesSign(r *mux.Router) {

	r.HandleFunc("/login", controllers.UserLoginForm).Methods("GET", "POST")
	r.HandleFunc("/logout", controllers.UserLogout).Methods("GET")
}
