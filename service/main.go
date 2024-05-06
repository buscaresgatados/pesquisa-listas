package main

import (
	"fmt"
	"net/http"
	"os"
	"refugio/handlers"

	"github.com/gorilla/mux"
)

var (
	port string
)

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pessoa", handlers.GetPessoa).Methods(http.MethodGet, http.MethodOptions).Queries()
	http.Handle("/", router)

	fmt.Println("Listening on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}

	// sheetscraper.Scrape()
}
