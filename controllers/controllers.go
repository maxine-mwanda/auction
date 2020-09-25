package controllers

import (
	"auction-api/entities"
	"auction-api/middleware"
	"auction-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Users(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		user, err := models.FetchUser()
		if err != nil {
			fmt.Println("unable to fetch users")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		middleware.JsonResponse(writer, 200, user)
	} else {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("unable to read users")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		var user entities.NewUser
		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("unable to read user to struct")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		id, err := models.Create_user(entities.NewUser{})
		if err != nil {
			fmt.Println(err)
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		middleware.JsonResponse(writer, 200, id)
	}

}

/*func User(writer http.ResponseWriter, request *http.Request) {
	values := mux.Vars(request)
	user := values["user"]
	if user.Id == 0 {
		middleware.JsonResponse(writer, 404, nil)
		return
	}
	middleware.JsonResponse(writer, 200, user)
}*/

func Product(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		productID := request.URL.Query().Get("product_id")
		if productID == "" {
			middleware.JsonResponse(writer, http.StatusBadRequest, "product id not given")
			return
		}
		product, err := models.FetchProduct(productID)
		if err != nil {
			fmt.Println("unable to fetch products")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		middleware.JsonResponse(writer, 200, product)
	} else {
		_, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("unable to read product")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}

	var product entities.NewProduct
		err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("unable to read product to struct")
		middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
		return
	}
	id, err := models.Create_product(entities.NewProduct{})
	if err != nil {
		fmt.Println(err)
		middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
		return
	}
	middleware.JsonResponse(writer, 200, id)
}

}