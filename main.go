package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP server")
	
	http.HandleFunc("/hello", routes)

	http.ListenAndServe(":8080", nil)
}

func routes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Hello World!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}
}