package handler

import "fmt"

func Catch() {
    if r := recover(); r != nil {
        fmt.Println("Error occured", r)
    } 
}