package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Start)
	log.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func Start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}
