package main

import (
	"fmt"
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
		err = tmpl.ExecuteTemplate(w,"index",index_data)
		if err != nil { 
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"about",about_data)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"contact",nil)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

		if r.Method == http.MethodPost{
			details := ContactDetails{
				r.FormValue("email"),
				r.FormValue("subject"),
				r.FormValue("body"),
			}
			fmt.Println(details.Email,details.Subject,details.Body)
			err = tmpl.ExecuteTemplate(w,"contact",struct{Success bool}{true})
			if err != nil{
				http.Error(w,err.Error(),http.StatusInternalServerError)
			}

		}
	})
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w,"doc",nil)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	})
	fmt.Println(Hello)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
