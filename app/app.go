package app

import (
	"adv/config"
	"adv/routers"
	"fmt"
	"net/http"
)

func App() {
	server := http.Server{
		Addr: config.ADDR,
		Handler: routers.Router(),
	}
	fmt.Printf("Server Running http://%v\n", server.Addr)
	server.ListenAndServe()
}

func App1 () {
	server := http.Server{
		Addr: config.ADDR,
		Handler: routers.RouteFile(),
		
	}
	fmt.Printf("Server Running http://%v\n", server.Addr)
	server.ListenAndServe()
}