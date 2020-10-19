package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main()  {
	
	r := mux.NewRouter()
	
	r.HandleFunc("/hello",helloHandler).Methods("GET")
	r.HandleFunc("/goodbye",goodByeHandler).Methods("GET")

	http.Handle("/",r)

	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080",nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	
	 fmt.Fprint(w, "Hello World")

}

func goodByeHandler(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Fprint(w, "Good bye")

}

