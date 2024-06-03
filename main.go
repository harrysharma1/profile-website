package main

import (
	"html/template"
	"log"
	"net/http"

)



func main()  {
	tmpl, err := template.ParseGlob("templates/*")  
	if err != nil{
		log.Fatal(err)	
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"index",nil)
		if err != nil { 
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"about",nil)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"contact",nil)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":6969", nil))
}
