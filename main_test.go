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

	output, err := jsonapi.UnmarshalManyPayload(rr.Body, reflect.TypeOf(&Article{}))
	if err != nil {
		t.Errorf("Failed to Unmarshall JSON response: ", err)
	}

	if len(output) != 2 {
		t.Fatalf("Expected `articles` reponse to contain two articles")
	}

	articles := make(Articles, len(output))
	for i, v := range output {
		fmt.Println("%#v", output)
		var ok bool
		if output[i], ok = v.(*Article); !ok {
			t.Errorf("%v is not an *Artilcle")
		}
	}

	fmt.Println("%#v", articles)
}
