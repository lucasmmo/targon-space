package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello World")
	})

	fmt.Printf("Server is running at %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}