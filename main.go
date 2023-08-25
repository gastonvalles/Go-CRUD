package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Start)
	log.PrintIn("Server running...")
}

func Start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}
