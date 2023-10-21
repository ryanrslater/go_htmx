package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("you have requested the file")

		tmpl := template.Must(template.ParseFiles("./index.html"))

		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)
}
