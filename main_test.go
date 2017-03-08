package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

	var output Articles
	jsonapi.UnmarshalPayload(rr.Body, &output)
	// articles := jsonapi.UnmarshalManyPayload(rr.Body, reflect.TypeOf(new(Article)))
	// fmt.Printf("%#v", articles)
	// TODO: Figure out why we are not unmarshaling the articles correctly

	fmt.Printf("%#v", output)
}
