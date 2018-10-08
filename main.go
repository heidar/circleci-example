package main

import (
  "net/http"
)

func RootHandler (res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Hello world!"))
}

func main() {
  http.HandleFunc("/", RootHandler)
  http.ListenAndServe(":3000", nil)
}
