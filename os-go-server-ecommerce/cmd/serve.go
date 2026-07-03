package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux() //router --> mux

	// route
	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on: 5000")
	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
