package main

import (
	"html/template"
	"log"
	"net/http"
)



func main()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/index.html"))		
		tmpl.Execute(w, nil)
	})	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/about.html"))
		tmpl.Execute(w,nil)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/contact.html"))
		tmpl.Execute(w,nil)
	})

	log.Fatal(http.ListenAndServe(":6969", nil))
}
