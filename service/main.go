package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"refugio/sheetscraper"
	"refugio/web"
	"refugio/web/handlers"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	port string
)

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
		router.Use(web.BaseRequestMiddleware)
		/* /pessoa routes with caching and Auth */
		pessoaSubrouter := router.PathPrefix("/pessoa").Subrouter()
		pessoaSubrouter.Use(web.AuthMiddleware, web.CacheMiddleware)
		pessoaSubrouter.HandleFunc("", handlers.GetPessoa).Methods(http.MethodGet, http.MethodOptions).Queries("nome", "{nome:[a-zA-Z0-9]{3,}}")
		pessoaSubrouter.HandleFunc("/count", handlers.GetRecordCount).Methods(http.MethodGet, http.MethodOptions)
		pessoaSubrouter.HandleFunc("/most_recent", handlers.GetMostRecent).Methods(http.MethodGet, http.MethodOptions)

		router.Handle("/sources", web.AuthMiddleware(http.HandlerFunc(handlers.GetSources))).Methods(http.MethodGet, http.MethodOptions)
		router.Handle("/auth/me", web.AuthMiddleware(http.HandlerFunc(handlers.AuthMe))).Methods(http.MethodGet, http.MethodOptions)

		router.HandleFunc("/health/ready", handlers.Ready).Methods(http.MethodGet, http.MethodOptions)
		router.HandleFunc("/health/live", handlers.Live).Methods(http.MethodGet, http.MethodOptions)

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
