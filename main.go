package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, req *http.Request){
	fmt.Println("Endpoint Hit: Return all Articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, req *http.Request){
	fmt.Println("Endpoint Hit: Return single Article")
	vars := mux.Vars(req)

	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)

	for _, article := range Articles{
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}		
	}
	
}

func createNewArticle(w http.ResponseWriter, req *http.Request){
	fmt.Println("Hit endpoint: Create New")
	reqBody, _ := io.ReadAll(req.Body)

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticleById(w http.ResponseWriter, req *http.Request){
	fmt.Println("Hit endpoint: Delete By Id")

	vars := mux.Vars(req)
	fmt.Println(vars)

	key := vars["id"]

	for index , article := range Articles{
		if article.Id == key{
			Articles = append(Articles[:index], Articles[index + 1:]...)
		}
	}
}

func updateArticleById(w http.ResponseWriter, req *http.Request){
	fmt.Println("Hit endpoint: Update Article")
	vars := mux.Vars(req)
	reqBody, _ := io.ReadAll(req.Body)
	key := vars["id"]

	var article Article 

	json.Unmarshal(reqBody, &article)

	for index, article := range Articles{
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index + 1:]...)
			Articles = append(Articles, article)
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
// ###### http Implementation #######
	// http.HandleFunc("/", homepage)
	// http.HandleFunc("/articles", returnAllArticles)

// ####### mux Implementation ########
	//creates a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticleById).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticleById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main(){
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc:"Article Description", Content:"Article Content"},
		Article{Id: "2", Title: "Hello2", Desc:"Article Description2", Content:"Article Content2"},
	}
	fmt.Println("Rest API 2.0 = Mux Router")
	fmt.Println("application started on port: http://localhost:8000")
	handleRequests()
}

