package handlers

import (
	"encoding/json"
	"net/http"
	"refugio/repository"
)

func Ready(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func Live(w http.ResponseWriter, r *http.Request) {
	results, _ := repository.FetchPessoaFromFirestore([]string{"yusbelydayarithfloresmaicanginsiopoliesportivolasalle"})
	if len(results) == 0 {
		http.Error(w, "Error fetching people from Firestore", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("OK"))
}

func AuthMe(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "success", "message": "Authenticated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
