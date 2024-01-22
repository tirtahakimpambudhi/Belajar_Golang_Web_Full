package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"
)

var secretKey = "secret"
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = secretKey
	cookie.Value = "token=aihe8iohbido23hdhb9ned8809du90di"
	cookie.Path = "/"

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

func TestSetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/login",nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder,request)

	cookies := recorder.Result().Cookies()

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)

	for _, c := range cookies {
		fmt.Printf("%v : %v\n",c.Name,c.Value)
	}

	fmt.Printf("\nResponse Body : %v\n",string(body))
}


func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8989",nil)
	recoder := httptest.NewRecorder()
	cookie := new(http.Cookie)
	cookie.Name = secretKey
	cookie.Value = "this message so secret"
	request.AddCookie(cookie)
	GetCookie(recoder,request)

	response := recoder.Result()
	body , _ := io.ReadAll(response.Body)

	fmt.Printf("\nResponse Body : %v\n",string(body))
}