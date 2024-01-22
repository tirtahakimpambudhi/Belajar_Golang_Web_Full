package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	_ "strings"
	"testing"
)

func helloHandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w,"hello")
}

func TestHttpTes(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/hello",nil)
	recoder := httptest.NewRecorder()

	helloHandler(recoder,request)

	response := recoder.Result()
	body , _ := io.ReadAll(response.Body)
	
	fmt.Printf("Response Body : %v",string(body))
}

func QueryParameterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w,"Hello %v",name)
}

func TestQuery(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/hello",nil)
	recoder := httptest.NewRecorder()

	QueryParameterHandler(recoder,request)

	response := recoder.Result()
	body , _ := io.ReadAll(response.Body)
	
	fmt.Printf("Response Body : %v",string(body))
}


func QueryMultiple(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	
	fmt.Fprintf(w,"\nID : (%v)\nName : '%v'",id,name)
}
func TestQueryMultiple(t *testing.T){
		request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/hello?id=1&name=tirta",nil)
		recoder := httptest.NewRecorder()
	
		QueryMultiple(recoder,request)
	
		response := recoder.Result()
		body , _ := io.ReadAll(response.Body)
		
		fmt.Printf("Response Body : %v",string(body))
}


func QueryMultipleValue(w http.ResponseWriter, r *http.Request) {
	var firstname, middlename, lastname string

	params := r.URL.Query()
	names, exists := params["name"]

	if exists {
		if names != nil && len(names) > 0 {
			// Check the underlying type using a type assertion
			if reflect.TypeOf(names).Kind() == reflect.Slice && reflect.TypeOf(names).Elem().Kind() == reflect.String {
				firstname = names[0]
				lastname = names[len(names)-1]
				for _, name := range names {
					if name != firstname && name != lastname {
						middlename += " " + name
					}
				}

				fmt.Fprintf(w, "\nFull Name :%v\nFirst Name : %v\nMiddle Name : %v\nLast Name : %v", strings.Join(names, " "), firstname, middlename, lastname)
				return
			}
		}
	}

	fmt.Fprintf(w, "\nFull Name : %v", params.Get("name"))
}


func TestQueryMultiValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/hello?name=tirta",nil)
	recoder := httptest.NewRecorder()

	QueryMultipleValue(recoder,request)

	response := recoder.Result()
	body , _ := io.ReadAll(response.Body)
	
	fmt.Printf("Response Body : %v",string(body))
}


func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	if contentType == "application/json"{
		message := map[string]string{
			"Content-Type":contentType,
		}
		messageJSON , err := json.Marshal(message)
		if err != nil {
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		}

		w.Header().Set("content-type",contentType)
		w.Write(messageJSON)
		return
	}

	if contentType == "text/html" {
		htmlMessage := fmt.Sprintf("<html><body><h1>Content-Type : <span>%v</span></h1></body></html>",contentType)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlMessage))
	}
	fmt.Fprintf(w,"\nContent-Type : %v",contentType)
}

//Client ==> content-type ==> server ==> message with content-type request client ==> Client
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/",nil)
	recoder := httptest.NewRecorder()
	request.Header.Add("Content-Type","application/json")
	RequestHeader(recoder,request)

	response := recoder.Result()
	body , _ := io.ReadAll(response.Body)
	
	fmt.Printf("Response Body : %v",string(body))
}
var key string = "x-power-by"
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(key,"secret")
	fmt.Fprint(w,"OK")
}
//Server ==> content-type ( set or add ) ==> Client get content-type
func TestResponseHeader(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet,"http://localhost:8888/",nil)
		recoder := httptest.NewRecorder()
		ResponseHeader(recoder,request)
		message := recoder.Header().Get(key)
		response := recoder.Result()
		body , _ := io.ReadAll(response.Body)
		
		fmt.Printf("Response Body : %v\nMessage : %v",string(body),message)
}


