package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct {
	Title string    'json : "Title"'
	Desc string     'json : "desc"'
	Content string  'json : "conent"'
}
type Articles []Article
func allArticles(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title:"TestTitle",Desc:"Test Description",Content:"Hello World"},
		fmt.Println("Endpoint Hit : All(ArticleEndpoint")
		json.NewEncoder(w).Encode(articles)
	}
func homepage(w http.ResponseWriter,r *http.Request{
	fmt.Fprintf(w, "Homepage End Hit")
}
func handleRequest(){
	myRouter := mux.NewRoute().StrictSlash(true)
	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}
})
func main(){
	handleRequest()
}
	}
}
}