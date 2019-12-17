package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Println("Hit me with a GET request")
	case "POST":
		fmt.Println("Hit me with a POST request")
	default:
		fmt.Println("Hit me with a request I don't recognize")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body", err)
	}

	fmt.Println(string(body))

	t, _ := template.ParseFiles("static/notes.html")
	t.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
