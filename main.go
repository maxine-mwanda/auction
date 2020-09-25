package main

import (
	"auction-api/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":5000"
	router := mux.NewRouter()
	fmt.Println("Hey there")

	router.HandleFunc("/users", controllers.Users).Methods("GET", "POST")
	router.HandleFunc("/user", controllers.User).Methods("GET")
	router.HandleFunc("/product", controllers.Product).Methods("GET")
	router.HandleFunc("/product", controllers.Product).Methods("POST")

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("unable to start api because ", err)
	}

}
