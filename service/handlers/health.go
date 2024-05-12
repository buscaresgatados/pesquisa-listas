package handlers

import (
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
