package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTfulll api")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
