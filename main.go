package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	fmt.Println("Starting server at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
