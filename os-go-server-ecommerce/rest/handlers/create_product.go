package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	// i need this when using --> HandleFunc
	// if r.Method != "POST" {
	// 	http.Error(w, "Please give me POST request", 400)
	// 	return
	// }

	/*
		1. take body information (description, imageUrl, price, title) from r.Body
		2. create an instance using Product struct with the body information
		3. append the instance into productList
	*/

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	// frontend theke json asbe oitake amra decode kore struct/js object e convert krbo.
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, 201)
}
