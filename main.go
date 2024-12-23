package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset-utf=8")
		err = tmpl.ExecuteTemplate(w, "index", index_data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset-utf=8")
		err = tmpl.ExecuteTemplate(w, "about", about_data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset-utf=8")
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			details := ContactDetails{
				Email:   r.FormValue("email"),
				Subject: r.FormValue("subject"),
				Body:    r.FormValue("message"),
			}
			fmt.Println(details)
			err = tmpl.ExecuteTemplate(w, "contact", struct{ Success bool }{true})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			err = tmpl.ExecuteTemplate(w, "contact", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
	http.HandleFunc("/cv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset-utf=8")
		err = tmpl.ExecuteTemplate(w, "cv", cv_data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	fmt.Println(Hello)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
