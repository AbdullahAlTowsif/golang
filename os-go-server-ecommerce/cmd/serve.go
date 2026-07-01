package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router --> mux

	// route
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on: 5000")
	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
