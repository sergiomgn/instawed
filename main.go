package main

import (
	"discardcam/models"
	"log"
	"net/http"
	"text/template"

	"github.com/golang-jwt/jwt/v4"
)

var tmpl = template.Must(template.ParseFiles("static/login.html"))

var jwtKey = []byte("secret_key") // This key should be changed only before runtime to avoid leaks

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

	user, err = models.FindUserByUsername(username)
	if err != nil {
		user = &models.User{
			Username: username,
			Email:    email,
		}
		err = models.CreateUser(user)
	}
	err = models.CreateUser(user)
	if err != nil {
		http.Error(w, "Error Creating User", http.StatusInternalServerError)
		return
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
