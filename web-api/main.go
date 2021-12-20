package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Jan!</h1>")
}

func main() {
	http.HandleFunc("/", helloPage)
	fmt.Println("Server running on 8080....")
	fmt.Println("Press Ctrl+C To STOP")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
