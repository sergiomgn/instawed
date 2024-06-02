package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("static/login.html"))

var users = map[string]string{
	"user": "email",
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")

	if storedEmail, ok := users[username]; ok && storedEmail == email {
		fmt.Fprintf(w, "Utilizador Registado!")
	} else {
		fmt.Fprintf(w, "Utilizador nao valido")
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	log.Print("Starting server on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
