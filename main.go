package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

var templates = template.New("app")

func main()  {

	templates = template.Must(template.ParseGlob("templates/*.html"))
	
	r := mux.NewRouter()
	
	r.HandleFunc("/",indexHandler).Methods("GET")


	http.Handle("/",r)
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080",nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	
	templates.ExecuteTemplate(w,"index.html",nil)

}


