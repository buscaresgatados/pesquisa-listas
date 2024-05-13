package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"refugio/objects"
	"strings"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"

	"cloud.google.com/go/compute/metadata"
)

var (
	authKeys []string
	cache    *expirable.LRU[string, interface{}]
)

func BaseRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "private")
		if r.Method == http.MethodOptions {
			return
		}

		if os.Getenv("ENVIRONMENT") != "local" {
			// Trace logging for Cloud Run
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
		}
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Preflight has no Authorization header
		if r.Method == http.MethodOptions || os.Getenv("ENVIRONMENT") == "local" {
			next.ServeHTTP(w, r)
			return
		}
		key := r.Header.Get("Authorization")
		if !isValidKey(key) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/* Caching */
type ResponseCapture struct {
	http.ResponseWriter
	Body bytes.Buffer
}

func (w *ResponseCapture) Write(data []byte) (int, error) {
	w.Body.Write(data) // Capture the response body
	return w.ResponseWriter.Write(data)
}

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		nome := r.URL.Query().Get("nome")
		var cacheKey string = nome
		if cacheKey == "" {
			cacheKey = r.URL.Path
		}

		if result, ok := cache.Get(cacheKey); ok {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonBytes)
			return
		}

		capture := &ResponseCapture{ResponseWriter: w}
		next.ServeHTTP(capture, r)

		var result interface{}
		if err := json.Unmarshal(capture.Body.Bytes(), &result); err == nil {
			cache.Add(cacheKey, result)
		} else {
			fmt.Fprintf(os.Stderr, "Error unmarshalling response %v", err)
		}
	})
}

func isValidKey(key string) bool {
	for _, validKey := range authKeys {
		if key == validKey {
			return true
		}
	}
	return false
}

func init() {
	keyFile, err := os.ReadFile(os.Getenv("AUTH_KEYS_FILE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading auth keys file %v", err)
	}
	var keys map[string]string

	err = json.Unmarshal(keyFile, &keys)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling auth keys %v", err)
	}
	for _, v := range keys {
		authKeys = append(authKeys, v)
	}

	cache = expirable.NewLRU[string, interface{}](5000, nil, time.Minute*60)
}
