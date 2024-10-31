package main

import (
	"fmt"
	"github.com/ifere/algotrivia/config/db"
	"github.com/ifere/algotrivia/config/middleware"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// registers each service and call the services start method
// each service requires a connection to mongodb
// each start method requires the instance of the app handler
// TODO refactor register service to accept various db connection tyes
func registerServices(router *mux.Router, conn dbs.MongoDB) {

	InitializeQuestions(conn).Start(router)

}

//func setupBugSnag() {
//	bugsnag.Configure(bugsnag.Configuration{
//		APIKey:       "",
//		ReleaseStage: os.Getenv("env"),
//		// The import paths for the Go packages containing your source files
//		ProjectPackages: []string{"main", "algotrivia"},
//	})
//
//}

func registerMiddleWares(router *mux.Router) *negroni.Negroni {
	logger()
	n := negroni.Classic()
	n.Use(middleware.Cors())
	n.UseHandler(router)
	return n
}

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	port := os.Getenv("PORT")
	fmt.Println("port",port)
	router := mux.NewRouter()
	var db dbs.MongoDB
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("db connected")

	n := registerMiddleWares(router)
	registerServices(router, db)

	//setupBugSnag()

	log.Printf("server running on port %v", port)

	err = http.ListenAndServe(":"+port, n)

	fmt.Println(err)

}

func logger() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.SetOutput(logger.Writer())
}
