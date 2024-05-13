package web

import (
	"bytes"
	"context"
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
	authKeys map[string]string
	cache    *expirable.LRU[string, interface{}]
)

type contextKey string

const (
	ACCESS_LOG_CONTEXT_KEY contextKey = "access_log"
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
		log := &objects.AccessLog{}

		ctx := context.WithValue(r.Context(), ACCESS_LOG_CONTEXT_KEY, log)
		r = r.WithContext(ctx)
		if os.Getenv("ENVIRONMENT") != "local" {
			trace := getTrace(r)
			userIp := r.Header.Get("X-Forwarded-For")

			log.Trace = &trace
			log.UserIP = &userIp
		}

		next.ServeHTTP(w, r)

		accessLog := ctx.Value(ACCESS_LOG_CONTEXT_KEY).(*objects.AccessLog)
		logJson, err := json.Marshal(accessLog)

		if *accessLog.Trace == "" {
			return
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Fprintln(os.Stdout, string(logJson))
		}
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
		// If key is valid, add user to AccessLog context
		if user, authorized := isValidKey(key); !authorized {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			ctx := r.Context()
			accessLog := ctx.Value(ACCESS_LOG_CONTEXT_KEY)
			if accessLog != nil {
				accessLog.(*objects.AccessLog).KeyUser = user
			}
			r = r.WithContext(ctx)
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

func isValidKey(key string) (*string, bool) {
	for validUser, validKey := range authKeys {
		if key == validKey {
			return &validUser, true
		}
	}
	return nil, false
}

func init() {
	keyFile, err := os.ReadFile(os.Getenv("AUTH_KEYS_FILE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading auth keys file %v", err)
	}

	err = json.Unmarshal(keyFile, &authKeys)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling auth keys %v", err)
	}

	cache = expirable.NewLRU[string, interface{}](50000, nil, time.Minute*60)
}

func getTrace(r *http.Request) string {
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
	return trace
}
