package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRootHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/", nil)

  if err != nil {
    t.Fatal(err)
  }

  rr      := httptest.NewRecorder()
  handler := http.HandlerFunc(RootHandler)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: %d, expected: %d", status, http.StatusOK)
  }

  expected := "Hello world!"
  if rr.Body.String() != expected {
    t.Error("handler returned unexpected body: ", rr.Body.String(), expected)
  }
}
