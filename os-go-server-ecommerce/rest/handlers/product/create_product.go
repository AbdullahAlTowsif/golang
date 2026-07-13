package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	// frontend theke json asbe oitake amra decode kore struct/js object e convert krbo.
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgUrl: req.ImgUrl,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	util.SendData(w, createdProduct, http.StatusCreated)
}

/*
	1. take body information (description, imageUrl, price, title) from r.Body
	2. create an instance using Product struct with the body information
	3. append the instance into productList
*/
