package main

import (
	"API Auction/controllers"
	"fmt"
	"log"
	"net/http"
)

package main

import (
"api/controllers"
"fmt"
"github.com/gorilla/mux"
"log"
"net/http"
)

func main() {
	port := ":5000"
	router := mux.NewRouter()
	fmt.Println("Hey there")

	router.HandleFunc("/users", controllers.User).Methods("GET", "POST")
	router.HandleFunc("/product", controllers.Product).Methods("GET")
	router.HandleFunc("/user", controllers.Users).Methods("GET")

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("unable to start api because ", err)
	}

}