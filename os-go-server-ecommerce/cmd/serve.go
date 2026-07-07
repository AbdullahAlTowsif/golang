package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	
	mux := http.NewServeMux() //router --> mux
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	fmt.Println("Server running on: 5000")
	err := http.ListenAndServe(":5000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
