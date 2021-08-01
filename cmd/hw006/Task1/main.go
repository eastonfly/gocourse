package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Res struct {
	Host        string              `json:"host"`
	User_agent  string              `json:"user_agent"`
	Request_uri string              `json:"request_uri"`
	Headers     map[string][]string `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	req := Res{r.Host, r.UserAgent(), r.RequestURI, r.Header}

	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s\n", data)
}

func main() {
	fmt.Println("Expecting connection on port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
