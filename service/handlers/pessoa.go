package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"refugio/repository"
	"refugio/utils/cuckoo"
	"sort"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "private")
	if r.Method == http.MethodOptions {
		return
	}
	projectID, _ := metadata.ProjectID()

	var trace string
	if projectID != "" {
		traceHeader := r.Header.Get("X-Cloud-Trace-Context")
		traceParts := strings.Split(traceHeader, "/")
		if len(traceParts) > 0 && len(traceParts[0]) > 0 {
			trace = fmt.Sprintf("projects/%s/traces/%s", projectID, traceParts[0])
		}
	}

	log := &objects.AccessLog{
		Trace:  trace,
		UserIP: r.Header.Get("X-Forwarded-For"),
	}
	logJson, err := json.Marshal(log)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, string(logJson))
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
	pessoas, err := repository.FetchPessoaFromFirestore(docIDs)
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
		if _, ok := seen[person.Nome+*person.SheetId]; !ok {
			seen[person.Nome+*person.SheetId] = true
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
	most_recent, err := repository.FetchMostRecent(repository.PessoasAbrigos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching most recent: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var result objects.PessoaMostRecentResult
	result.Timestamp = most_recent
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)
}
