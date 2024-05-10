package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"refugio/repository"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

func GetSources(w http.ResponseWriter, r *http.Request){
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

	sources, err := repository.FetchSourcesFromFirestore()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonBytes, err := json.Marshal(sources)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}