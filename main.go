package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
)


var client *redis.Client
var templates = template.New("app")

func main()  {

	client = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
	})

	templates = template.Must(template.ParseGlob("templates/*.html"))
	
	r := mux.NewRouter()
	
	r.HandleFunc("/",indexHandler).Methods("GET")


	http.Handle("/",r)
	fmt.Println("Running on port 8088")
	http.ListenAndServe(":8088",nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	
	comments,err := client.LRange("comments",0,10).Result()

	if err != nil {
		return
	}
	templates.ExecuteTemplate(w,"index.html",comments)

}


