package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"refugio/repository"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "private")
	if r.Method == http.MethodOptions {
		return
	}
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		http.Error(w, "nome é obrigatório", http.StatusBadRequest)
		return
	}

	client := search.NewClient(os.Getenv("ALGOLIA_CLIENT"), os.Getenv("ALGOLIA_API_KEY"))
	index := client.InitIndex(os.Getenv("ALGOLIA_INDEX"))

	results, err := index.Search(nome)
	if err != nil {
		panic(err)
	}

	var pessoasSearch []objects.PessoaSearchResult
	err = results.UnmarshalHits(&pessoasSearch)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var pessoas []*objects.PessoaResult
	for _, pessoa := range pessoasSearch {
		pessoaRepository, err := repository.FetchFromFirestore(pessoa.ObjectID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pessoas = append(pessoas, pessoaRepository)
	}

	jsonBytes, err := json.Marshal(pessoas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
