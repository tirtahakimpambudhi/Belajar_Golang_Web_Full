package test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateString(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>{{.}}</body></html>`
	t , err := template.New("helloworld").Parse(html)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}

	t.ExecuteTemplate(w,"helloworld","Hello World")
}


func TestTemplateString(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:6969",nil)
	recorder := httptest.NewRecorder()

	TemplateString(recorder,request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}