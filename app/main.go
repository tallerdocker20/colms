package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
	"github.com/202lp1/colms/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("init")
}

var err error

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World!")
	//})
	fmt.Println("main")

	cfig.DB, err = connectDBmysql()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Printf("db is connected: %v", cfig.DB)

	// Migrate the schema
	cfig.DB.AutoMigrate(
		&models.Empleado{},
		&models.Alumno{},
		&models.Matricula{},

		&models.User{})
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	r := NewRouter()

	routes.RoutesSign(r)
	routes.RoutesMain(r)

	port := os.Getenv("PORT")
	log.Printf("porti: %v", port)
	if port == "" {
		port = "8080"
	}
	log.Printf("port: %v", port)
	http.ListenAndServe(":"+port, r)

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Choose the folder to serve
	staticDir := "/assets/"
	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	return router
}

func connectDBmysql() (c *gorm.DB, err error) {
	dsn := os.Getenv("MYSQL_DNS_DOKER")
	//dsn := os.Getenv("MYSQL_DNS_LOCAL1")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return conn, err
}

func connectDB() (c *gorm.DB, err error) {
	dsn := os.Getenv("PG_DNS_HEROKU")
	//dsn := os.Getenv("PG_DNS_LOCAL1")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return conn, err
}
