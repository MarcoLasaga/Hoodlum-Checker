package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	loadTemplates()

	mux := http.NewServeMux()

	mux.HandleFunc("/", landingHandler)
	mux.HandleFunc("/quiz", quizHandler)
	mux.HandleFunc("/result", resultHandler)

	fmt.Println("Running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
