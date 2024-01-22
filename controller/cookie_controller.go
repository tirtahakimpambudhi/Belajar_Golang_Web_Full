package controller

import (
	"fmt"
	"net/http"
)

var secretKey = "secret"
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = secretKey
	cookie.Value = "this message secret"
	cookie.Path = "/set-cookie"
	cookie.HttpOnly = true

	http.SetCookie(w,cookie)
	fmt.Fprint(w,"Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie , err := r.Cookie(secretKey)
	if err != nil {
		fmt.Fprint(w,"No Cookie")
		return
	}

	fmt.Fprintf(w,"Value Cookie : '%v'",cookie.Value)
}