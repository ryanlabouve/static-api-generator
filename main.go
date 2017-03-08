package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

type Article struct {
	ID          int       `jsonapi:"primary,articles"`
	Title       string    `jsonapi:"attr,title"`
	Description string    `jsonapi:"attr,description"`
	Date        time.Time `jsonapi:"attr,date"`
	Slug        string    `jsonapi:"attr,slug"`
	Content     string    `jsonapi:"attr,content"`
}

type Articles []*Article

var articles Articles

func init() {
	articles = Articles{
		&Article{
			1,
			"Test Title",
			"Here is my test description",
			time.Now(),
			"test-slug",
			"My content...",
		},
		&Article{
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
	fmt.Printf("Size of Articles: %d\n", len(articles))
	if err := jsonapi.MarshalManyPayload(w, articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	return router
}

func main() {
	// [x] : Goal one, list some articles
	// [x] : Goal two, list some articles in JSON API format
	// [ ] : Goal three, Add some tests using standard test package
	// [ ] : Goal x, filter QP's
	// [ ] : Goal x, Connect Ember
	// [ ] : Goal x, Write search
	// [ ] : Goal x, Do markdown to memory import
	// [ ] : Goal x, Deploy and cut over?
	log.Fatal(http.ListenAndServe(":7111", Router()))
}
