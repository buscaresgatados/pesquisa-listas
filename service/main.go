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
	authKeys = []string{
		"c2585727-bd1d-4b70-bd97-b0417c8e3c7c", // frontend
		"hardcoded-key2",
		"hardcoded-key3",
	}
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		router.Handle("/auth/me", AuthMiddleware(http.HandlerFunc(AuthMeHandler))).Methods(http.MethodGet, http.MethodOptions)

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
	scraperCmd.Flags().Bool("isDryRun", false, "Enable dry-run mode without making actual changes")
}
