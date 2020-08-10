package main

import (
	"html/template"
	"log"
	"net/http"
    
	"github.com/preethimaliMalki/crud/Crud"
)



func main() {
	log.Println("Server started on: http://localhost:9092")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/Insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":9092", nil)
}
