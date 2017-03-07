package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
}

type Articles []Article

var articles Articles

func init() {
	articles = Articles{
		Article{
			"Test Title",
			"Here is my test description",
			"0001-01-01T00:00:00Z",
			"test-slug",
			"My content...",
		},
		Article{
			"Test Title 2",
			"Here is my test description 2",
			"0001-01-01T00:00:00Z",
			"test-slug2",
			"My other content...",
		},
	}
}

// /articles
// /articles?filter[limit]=5
// /articles?filter[slug]=search string
// /articles?filter[term]=search term

func GetArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func main() {
	// [x] : Goal one, list some articles
	router := mux.NewRouter()
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":7111", router))
}
