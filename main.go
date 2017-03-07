package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

type Article struct {
	ID          int       `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
}

type Articles []interface{}

var articles Articles

func init() {
	articles = Articles{
		Article{
			1,
			"Test Title",
			"Here is my test description",
			time.Now(),
			"test-slug",
			"My content...",
		},
		Article{
			2,
			"Test Title 2",
			"Here is my test description 2",
			time.Now(),
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalManyPayload(w, articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// [x] : Goal one, list some articles
	// [ ] : Goal two, list some articles in JSON API format
	router := mux.NewRouter()
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":7111", router))
}
