package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", "Saptarsi")
	if err != nil {
		log.Fatalln(err)
	}
}
