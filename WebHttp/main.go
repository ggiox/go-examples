package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article is data struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles is a Array of data struct
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Titles", Desc: "Teste Description", Content: "Hello World"},
		Article{Title: "Test Titles 1", Desc: "Teste Description 1", Content: "Hello World 1"},
		Article{Title: "Test Titles 2", Desc: "Teste Description 2", Content: "Hello World 2"},
	}

	fmt.Fprintln(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Http Server Listen port 8000")
	handleRequests()
}
