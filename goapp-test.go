package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var greeting = "Hello World!"

//TestIndexHandler check if the webservice is giving the greeting value inside
func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := greeting
	if rr.Body.String() != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			rr.Body.String(),
			greeting,
		)
	}
}

//TestIndexHandlerNotFound a test to get 404 errors
func TestIndexHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/404", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusNotFound,
		)
	}
}
