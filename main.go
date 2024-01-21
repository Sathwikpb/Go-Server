package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST  req succeesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/hello" {
	// 	http.Error(w, "404 not found ", http.StatusNotFound)
	// 	return
	// }
	// if r.Method != "GET" {
	// 	http.Error(w, "method not supported", http.StatusNotFound)
	// 	return
	// }
	// fmt.Fprintf(w, "hello")
// Alternative method for path request 
	switch r.URL.Path {
	case "/hello":
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprint(w, "hello")
	default:
		http.NotFound(w, r)
	}

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting on 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
