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
	

		if r.Method == http.MethodPost{
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			details := ContactDetails{
				Email: r.FormValue("email"),
				Subject: r.FormValue("subject"),
				Body: r.FormValue("message"),
			}
			fmt.Println(details)
			err = tmpl.ExecuteTemplate(w,"contact",struct{Success bool}{true})
			if err != nil{
				http.Error(w,err.Error(),http.StatusInternalServerError)
			}

		}else{
			err = tmpl.ExecuteTemplate(w,"contact",nil)
			if err != nil {
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
