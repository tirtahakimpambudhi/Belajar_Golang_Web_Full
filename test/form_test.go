package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	_"strconv"
	"strings"
	"testing"
)



func Form(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w,"Error Parse Form",http.StatusInternalServerError)
		return
	}

	body := r.PostForm
	test := make(map[string]string)
	errorMessage := make(map[string]string)

	for key , el := range body {
		if el[0] == "" {
			errorMessage["errors"] = fmt.Sprintf("cannot '%v' empty",key)
			break
		}
		test[key] = el[0]
	}
	if len(errorMessage) > 0 {
		errorsJSON , _ := json.Marshal(errorMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorsJSON)	
		return
	}
	testJSON, _ := json.Marshal(test)

    w.Header().Set("Content-Type", "application/json")
    w.Write(testJSON)

}

func TestPostFormSuccess(t *testing.T) {
	requestBody := strings.NewReader("id=1&username=example&email=example@gmail.com&password=12324556")
	request := httptest.NewRequest(http.MethodPost,"http://localhost:8888",requestBody)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	Form(recorder,request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)
	fmt.Printf("Response Body : %v\n",string(body))

	parsedBody := make(map[string]string)

    _ = json.Unmarshal(body, &parsedBody)

    // Iterate over the map and print key-value pairs
    for key, value := range parsedBody {
		fmt.Printf("\nKey: %v, Value: %v\n", key, value)
    }
}
func TestPostFormInvalid(t *testing.T) {
	requestBody := strings.NewReader("id=1&username=example&email=example@gmail.com&password=")
	request := httptest.NewRequest(http.MethodPost,"http://localhost:8888",requestBody)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	Form(recorder,request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)
	fmt.Printf("Status Code : %v\n",response.StatusCode)
	fmt.Printf("Response Body : %v\n",string(body))

	parsedBody := make(map[string]string)

    _ = json.Unmarshal(body, &parsedBody)

    // Iterate over the map and print key-value pairs
    for key, value := range parsedBody {
		fmt.Printf("\nKey: %v, Value: %v\n", key, value)
    }
}
