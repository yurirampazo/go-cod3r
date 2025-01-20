package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rightTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("19/01/2025 03:04:05")
	fmt.Fprintf(w, "<h1>Right time: %s</h1>", s)
}

func main() {
	http.HandleFunc("/rightTime", rightTime)
	log.Println("Excecuting...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}