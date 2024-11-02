package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ifere/algotrivia/config/db"
	"github.com/ifere/algotrivia/config/middleware"
	"github.com/ifere/algotrivia/src/api"
	// "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

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
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Println("Could not load Env Variables")
	// 	fmt.Println(e)

	// }
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }
	fmt.Println("port",port)
	router := mux.NewRouter()
	var db dbs.MongoDB
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MONGO_DEV_URL:", os.Getenv("MONGO_DEV_URL"))
	fmt.Println("db connected")

	n := registerMiddleWares(router)
	registerServices(router, db)
	api.RegisterRoutes(router)

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
