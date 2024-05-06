package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"refugio/repository"
	"sort"

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

	docIDs := make([]string, 0, len(pessoasSearch))
	for _, result := range pessoasSearch {
		docIDs = append(docIDs, result.ObjectID)
	}
	pessoas, err := repository.FetchFromFirestore(docIDs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sort.SliceStable(pessoas, func(i, j int) bool {
		return pessoas[i].Timestamp.After(pessoas[j].Timestamp)
	})

	// Deduplicate by Pessoa.Nome + Pessoa.SheetId
	unique := make([]*objects.PessoaResult, 0)
	seen := make(map[string]bool)

	for _, person := range pessoas {
		if _, ok := seen[person.Nome+person.SheetId]; !ok {
			seen[person.Nome+person.SheetId] = true
			unique = append(unique, person)
		}
	}

	jsonBytes, err := json.Marshal(unique)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
