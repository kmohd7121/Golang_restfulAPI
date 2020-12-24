package main

import ("fmt"
	"log"
	"net/http")
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Homepage Endpoint Hit")
}
func handleRequestds(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main(){
	handleRequestds()
}