package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}

	err := server.ListenAndServe() 
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}