package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomePage struct {
	Title string
	Shows [4]Show
}

type Show struct {
	Title       string
	Description string
	Author      string
	Date        string
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("you have requested the file")

		ShowList := [4]Show{
			{Title: "Show 1", Description: "Lorem ipsum", Author: "Ryan", Date: "28/2/21"},
			{Title: "Show 2", Description: "Lorem", Author: "Ryan", Date: "28/4/21"},
			{Title: "Show 3", Description: "Ip", Author: "Rryan", Date: "31/2/21"},
			{Title: "Show 4", Description: "fasdfasdf", Author: "Ryan", Date: "5/12/21"},
		}

		HomePageData := HomePage{
			Title: "Upcoming Show",
			Shows: ShowList,
		}

		tmpl := template.Must(template.ParseFiles("./templates/index.html"))

		tmpl.Execute(w, HomePageData)
	})

	http.ListenAndServe(":80", nil)
}
