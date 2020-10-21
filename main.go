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
	
	r.HandleFunc("/",indexGetHandler).Methods("GET")
	r.HandleFunc("/",indexPostHandler).Methods("POST")

	//File Server
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))

	
	http.Handle("/",r)
	fmt.Println("Running on port 8088")
	http.ListenAndServe(":8088",nil)

}

func indexGetHandler(w http.ResponseWriter, r *http.Request)  {
	
	comments,err := client.LRange("comments",0,10).Result()

	if err != nil {
		return
	}
	templates.ExecuteTemplate(w,"index.html",comments)

}

func indexPostHandler(w http.ResponseWriter, r *http.Request)  {
	
	r.ParseForm()

	comment := r.PostForm.Get("comment")
	client.LPush("comments",comment)

	http.Redirect(w,r,"/",302)
}


