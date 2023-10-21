package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Println(w, "you have requestd", r.URL.Path)
  })
  http.ListenAndServe(":80", nil)
} 
