package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hello, world!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("public/index.html")
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(body))
}
