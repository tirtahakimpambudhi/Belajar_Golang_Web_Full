package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func Form(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	body := r.PostForm
	values := make(map[string]string)
	errorsMessage := make(map[string]string)

	if len(body) == 0 {
		errorsMessage["errors"] = "BAD REQUEST"
	} else {
		for key, element := range body {
			if element[0] == "" {
				errorsMessage["errors"] = fmt.Sprintf("field '%v' cannot empty",key)
				break
			}
			values[key] = element[0]
		}
	}

	if len(errorsMessage) > 0 {
		errorsJSON , _ := json.Marshal(errorsMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("content-type","application/json")
		w.Write([]byte(errorsJSON))
		return
	}
	resultJSON , err := json.Marshal(values)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	w.Header().Set("content-type","application/json")
	w.Write([]byte(resultJSON))
}