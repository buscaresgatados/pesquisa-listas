package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"refugio/repository"
	"refugio/utils/cuckoo"
)

func GetRecordCount(w http.ResponseWriter, r *http.Request) {
	filter, err := cuckoo.GetCuckooFilter(repository.PessoasAbrigos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting filter: %v\n", err)
	}

	var result objects.PessoaCountResult
	result.Total = int(filter.Count())
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func GetMostRecent(w http.ResponseWriter, r *http.Request) {
	result, err := repository.FetchMostRecent(repository.PessoasAbrigos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching most recent: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
}
