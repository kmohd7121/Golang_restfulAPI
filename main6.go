package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// post is a struct that re
type Post struct{
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
} 
var data []Post= []Post{}
func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/add",addItem).Methods("POST")
	http.ListenAndServe(":5000",router)
}

func addItem(w http.ResponseWriter,r *http.Request){
	// routeVariable:=mux.Vars(r)["item"]
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	data = append(data,newPost)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}