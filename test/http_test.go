package test

import (
	"adv/config"
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var mainHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr: config.ADDR,
		Handler: mainHandler,
	}
	t.Logf("Server Running http://%s",server.Addr)
	server.ListenAndServe()
}
func HelloController(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w, "Hello World")
}
func TestServerMux(t *testing.T){
	router := http.NewServeMux()
	router.HandleFunc("/",HelloController)

	server := http.Server{
		Addr: config.ADDR,
		Handler: router,
	}
	t.Logf("Server Running http://%s",server.Addr)
	server.ListenAndServe()
}


