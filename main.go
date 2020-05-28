package main

import (
	"log"
	"net/http"
	"os"
	"sprint3/Helpers"
	"sprint3/Routers"
	_ "sprint3/docs"

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Templates Manager API
// @version 1.0
// @description This is a sample service for generating templates
// @termsOfService http://swagger.io/terms/
// @contact.name Essami nader
// @contact.email nader@craftfoundry.tech
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host backendtmpl.herokuapp.com
// @BasePath /
func main() {

	Helpers.Migration()
	mux := Routers.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("Listening on port ", port)

	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://frontendtmpl.herokuapp.com/", "http://localhost:3000"}))(mux))
	if err != nil {
		panic(err)

	}

}
