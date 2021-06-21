## GO GORM CRUD

Realiza el crud de una tabla de base de datos con gorm

* [Resources sharing to docker](#resources-sharing-to-docker)
* [A. Runing local](#a.-runing-local)
* [B. Runing form Docker](#b.-runing-form-Docker)
* [Deploy en heroku](#deploy-en-heroku)
* [License](#license)


## Resources sharing to docker

	Add D:\dockr

Dockerfile
```bash
FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /app/server
COPY go.mod .
COPY go.sum .

RUN go get github.com/cespare/reflex

RUN go mod download
COPY . .

RUN go build 
# CMD ["./server"]

#CMD make watch
```

docker-compose.yml
```bash
version: '3'

services:
  colms:
    build:
      context: "./app"
    volumes:
      - "./app:/app/server"
    container_name: colms-app
    restart: always
    ports:
      - "8090:8080"
    tty: true

#para la db unirse a una red de docker
networks:
  default:
    external:
      name: mysql_default # deberá estar corriendo este contenedor.  docker network ls

```

### A. Runing local  

```bash
PS D:\dockr\lp1\colms\app> nodemon --exec go run main.go --signal SIGTERM

```


### B. Runing form Docker

Build docker project

```bash
PS D:\dockr\lp1\colms> docker-compose up --build -d
PS D:\dockr\lp1\colms> docker ps
CONTAINER ID        IMAGE                         COMMAND                  CREATED             STATUS              PORTS                    NAMES
36b836b4c783        colms_colms                   "bash"                   8 minutes ago       Up 7 minutes        0.0.0.0:8090->8080/tcp   colms-app


PS D:\dockr\lp1\colms> docker exec -it colms-app bash

or
PS D:\dockr\lp1\colms> docker exec -it colms-app sh

```

Running

```bash
PS D:\dockr\lp1\colms> docker exec -it colms-app bash

root@22be69ba019e:/app/server# make watch
```
Ir a http://localhost:8090

#### NOTA: La base de datos esta en heroku aws
go.mod dependencias del proyecto

```bash
module github.com/202lp1/colms

// +heroku goVersion go1.14
go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/twinj/uuid v1.0.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.7
)
```

main.go code

```bash
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/controllers"
	"github.com/202lp1/colms/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {

	cfig.DB, err = connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Printf("db is connected: %v", cfig.DB)
	
	// Migrate the schema
	cfig.DB.AutoMigrate(&models.Empleado{})
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/item/index", controllers.ItemList).Methods("GET")

	r.HandleFunc("/employee/index", controllers.EmployeeList).Methods("GET")
	r.HandleFunc("/employee/form", controllers.EmployeeForm).Methods("GET", "POST")
	r.HandleFunc("/employee/delete", controllers.EmployeeDel).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
	  port = "8080"
	}
	log.Printf("port: %v", port)
	http.ListenAndServe(":"+port, r)

}

func connectDB() (c *gorm.DB, err error) {
	////dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "user=xhaahqndliodvx password=7158d25578fc8450d49e1cd2175c42eca6e25910fbe1588270500e2ecf47ee77 host=ec2-34-204-121-199.compute-1.amazonaws.com dbname=d4ta8dj9qr5u62 port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//dsn := "user=postgres password=postgres2 dbname=users_test host=localhost port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return conn, err
}
```

### Deploy en heroku

[colms-toheroku]:      https://github.com/202lp1/colms-toheroku

Please go to [colms-toheroku] es el proyecto para subir a heroku.
Copy the `app` files and paste into [colms-toheroku] root.


### License



GNU, see [LICENSE](LICENSE).

Equipo de investigación y desarrollo: 
- angeli@upeu.edu.pe, 
