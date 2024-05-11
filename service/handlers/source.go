package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/repository"
)

func GetSources(w http.ResponseWriter, r *http.Request) {
	sources, err := repository.FetchSourcesFromFirestore()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(sources)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
