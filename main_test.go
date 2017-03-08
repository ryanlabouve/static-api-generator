package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/jsonapi"
)

func TestGetArticles(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	if e, a := http.StatusOK, rr.Code; e != a {
		t.Fatalf("Expected a status of %d, got %d", e, a)
	}

	articles, err := jsonapi.UnmarshalManyPayload(rr.Body, reflect.TypeOf(new(Article)))
	if err != nil {
		fmt.Printf("%#v", err)
	}

	if len(articles) != 2 {
		t.Fatalf("Expected `articles` reponse to contain two articles")
	}
}
