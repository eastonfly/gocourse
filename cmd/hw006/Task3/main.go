package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("eastonfly")

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")

		token := jwt.New(jwt.SigningMethodHS256)

		claims := make(jwt.MapClaims)
		claims["name"] = name + ":" + address
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		token.Claims = claims

		tokenString, _ := token.SignedString(mySigningKey)

		cookie := &http.Cookie{
			Name:  "token",
			Value: tokenString,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
