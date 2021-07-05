package main

import (
	"fmt"
	"log"
	"net/http"
)

const webContent = "Hello World!"

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello2", helloHandler)
	log.Fatal(http.ListenAndServe(":9091", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
