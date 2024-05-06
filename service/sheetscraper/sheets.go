package sheetscraper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"refugio/objects"
	"refugio/repository"
	"refugio/utils"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	resultsSheetRange = "Bacon Log!A1:B1"
)

type SheetsSource struct{}
type SourceData struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
}

func (ss *SheetsSource) Read(sheetID string, sheetRange string) (interface{}, error) {
	serviceAccJSON := utils.GetServiceAccountJSON(os.Getenv("SHEETS_SERVICE_ACCOUNT_JSON"))
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(serviceAccJSON))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(sheetID, sheetRange).Do()
	if err != nil {
		return nil, err
	}

	return resp.Values, nil
}

func (ss *SheetsSource) LogStatus(sheetID string, status string) error {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(utils.GetServiceAccountJSON(os.Getenv("SHEETS_SERVICE_ACCOUNT_JSON"))))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	var valueRange sheets.ValueRange
	valueRange.Values = append(valueRange.Values, []interface{}{status, time.Now().Local().String()})
	_, err = srv.Spreadsheets.Values.Append(sheetID, resultsSheetRange, &valueRange).InsertDataOption("INSERT_ROWS").ValueInputOption("RAW").Do()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on spreadsheet %s: %v", sheetID, err)
		return nil
	}
	return nil
}

func Scrape(isDryRun bool) {
	ss := SheetsSource{}
	var serializedData []*objects.PessoaResult

	for _, cfg := range Config {
		for _, sheetRange := range cfg.sheetRanges {
			content, _ := ss.Read(cfg.id, sheetRange)
			fmt.Fprintf(os.Stdout, "Scraping data from sheetId %s, range %s", cfg.id, sheetRange)
			switch sheetRange {
			// Offsets e customizações pra cada planilha hardcoded por enquanto
			case "Alojados!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 10 || len(row) == 0 {
						continue
					}

					p := objects.Pessoa{
						Abrigo: row[2].(string),
						Nome:   row[3].(string),
					}

					if len(row) > 4 {
						p.Idade = row[4].(string)
					} else {
						p.Idade = ""
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CADASTRO_EM_TEMPO_REAL!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: row[2].(string),
						Nome:   row[1].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "ALOJADOS x ABRIGOS!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 13 || len(row) < 4 {
						continue
					}

					p := objects.Pessoa{
						Abrigo: row[2].(string),
						Nome:   row[3].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "PESSOAS RESGATADAS":
				for i, row := range content.([][]interface{}) {

					if i < 4 || len(row) < 3 {
						continue
					}

					p := objects.Pessoa{
						Abrigo: row[2].(string),
						Nome:   row[0].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			}
			if !isDryRun {
				repository.AddToFirestore(serializedData)
			}

			fmt.Fprintf(os.Stdout, "Scraped data from sheetId %s, range %s. %d results. Dry run? %v", cfg.id, sheetRange, len(serializedData), isDryRun)
		}
	}
}
