package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"refugio/handlers"
	"refugio/sheetscraper"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	port     string
	authKeys []string
)

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

func isValidKey(key string) bool {
	for _, validKey := range authKeys {
		if key == validKey {
			return true
		}
	}
	return false
}

func AuthMeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "success", "message": "Authenticated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(scraperCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		router := mux.NewRouter()

		router.Handle("/pessoa", AuthMiddleware(http.HandlerFunc(handlers.GetPessoa))).Methods(http.MethodGet, http.MethodOptions).Queries()
		router.Handle("/pessoa/count", AuthMiddleware(http.HandlerFunc(handlers.GetRecordCount))).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/pessoa/most_recent", AuthMiddleware(http.HandlerFunc(handlers.GetMostRecent))).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/sources", AuthMiddleware(http.HandlerFunc(handlers.GetSources))).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/auth/me", AuthMiddleware(http.HandlerFunc(AuthMeHandler))).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/health/ready", http.HandlerFunc(handlers.Ready)).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/health/live", http.HandlerFunc(handlers.Live)).Methods(http.MethodGet, http.MethodOptions)

		http.Handle("/", router)

		fmt.Println("Listening on port ", port)
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		if err != nil {
			panic(err)
		}
	},
}

var scraperCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Run the sheetscraper",
	Run: func(cmd *cobra.Command, args []string) {
		isDryRun, _ := cmd.Flags().GetBool("isDryRun")
		sheetscraper.Scrape(isDryRun)
	},
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
	scraperCmd.Flags().Bool("isDryRun", false, "Enable dry-run mode without making actual changes")
}
