package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// post is a struct that re
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

var data []Post = []Post{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	http.ListenAndServe(":5000", router)
}
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// routeVariable:=mux.Vars(r)["item"]
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
