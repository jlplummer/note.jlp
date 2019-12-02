package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Access-Control-Allow-Origin", "*")

	/*if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}*/

	fmt.Println("url path:", r.URL.Path)
	http.ServeFile(w, r, "notes.html")

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
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
