package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandleGetGreetRR tests the handler using an HTTP recorder.
func TestHandleGetGreetRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/greet", nil)
	if err != nil {
		t.Fatal(err)
	}

	handleGetGreet(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, but got %d", rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()

	expected := "Jay Shree Ram"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != expected {
		t.Errorf("Expected %s but got %s", expected, string(b))
	}
}

// TestHandleGetFoo tests the handler using an actual HTTP server.
func TestHandleGetGreet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetGreet))
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", response.StatusCode)
	}

	expected := "Jay Shree Ram"
	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != expected {
		t.Errorf("Expected %s but got %s", expected, string(b))
	}
}
