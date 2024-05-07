package sheetscraper

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
			content, err := ss.Read(cfg.id, sheetRange)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading sheet %s: %v", cfg.id, err)
				continue
			}
			fmt.Fprintf(os.Stdout, "Scraping data from sheetId %s, range %s", cfg.id, sheetRange)
			sheetNameAndRange := cfg.id + sheetRange
			switch sheetNameAndRange {
			// Offsets e customizações pra cada planilha hardcoded por enquanto
			case cfg.id + "Alojados!A1:ZZ":
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
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CADASTRO_EM_TEMPO_REAL!A1:ZZ":
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
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ALOJADOS x ABRIGOS!A1:ZZ":
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
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ATUALIZADO 06/05!A1:ZZ":
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
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ESCOLA ANDRÉ PUENTE!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 2 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Escola André Puente",
						Nome:   row[0].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "EMEF WALTER PERACCHI DE BARCELLOS!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 2 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "EMEF Walter Peracchi de Barcellos",
						Nome:   row[1].(string),
						Idade:  row[2].(string),
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CACHOEIRINHA!A1:ZZ":
				for _, row := range content.([][]interface{}) {
					p := objects.Pessoa{
						Abrigo: row[1].(string),
						Nome:   row[0].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "COLÉGIO MARIA AUXILIADORA!A1:ZZ":
				for _, row := range content.([][]interface{}) {
					p := objects.Pessoa{
						Abrigo: "Colégio Maria Auxiliadora",
						Nome:   row[0].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ULBRA - Prédio 14!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "ULBRA - Prédio 14",
						Nome:   row[0].(string),
					}

					if len(row) > 2 {
						p.Idade = row[1].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "COLÉGIO MIGUEL LAMPERT!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Colégio Miguel Lampert",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "AMORJI!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Associação dos Moradores do Jardim Igara II - AMORJI",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ESCOLA RONDONIA!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Escola Rondônia",
						Nome:   row[0].(string),
					}

					if len(row) > 2 {
						p.Idade = row[1].(string)
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Escola Jacob Longoni!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Escola Jacob Longoni",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "COLÉGIO ESPÍRITO SANTO!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}

					var p objects.Pessoa
					pattern := `^(.*?)(\d+)`
					re := regexp.MustCompile(pattern)
					matches := re.FindStringSubmatch(row[0].(string))
					if len(matches) > 0 {
						p = objects.Pessoa{
							Abrigo: "Colégio Espirito Santo",
							Nome:   matches[1],
							Idade:  matches[2],
						}
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CLUBE DOS EMPREGADOS DA PETROBRÁS!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}

					split := strings.Split(row[0].(string), "\n")
					for _, s := range split {
						p := objects.Pessoa{
							Abrigo: "Clube dos Empregados da Petrobras",
							Nome:   s,
							Idade:  "",
						}

						fmt.Fprintln(os.Stdout, p)
						serializedData = append(serializedData, &objects.PessoaResult{
							Pessoa:    &p,
							SheetId:   &cfg.id,
							Timestamp: time.Now(),
						})
					}
				}
			case cfg.id + "Colegio Guajuviras!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 2 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Colégio Guajuviras",
						Nome:   row[0].(string),
					}

					if len(row) > 2 {
						p.Idade = row[1].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CEL São José!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "CEL São José",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CR BRASIL!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "CR Brasil",
						Nome:   row[2].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CSSGAPA!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Associação de Suboficiais e Sargentos da Guarnição de Aeronáutica de Porto Alegre",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CTG Brazão do Rio Grande!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "CTG Brazão do Rio Grande",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CTG Seiva Nativa!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "CTG Seiva Nativa",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "EMEF ILDO!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "EMEF Ildo Meneghetti",
						Nome:   row[1].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Escola Irmao pedro!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Escola Irmão Pedro",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "FENIX!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Abrigo Fenix",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "PARÓQUIA SANTA LUZIA!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}

					split := strings.Split(row[0].(string), "\n")
					for _, s := range split {
						p := objects.Pessoa{
							Abrigo: "Paróquia Santa Luzia",
							Nome:   s,
							Idade:  "",
						}

						fmt.Fprintln(os.Stdout, p)
						serializedData = append(serializedData, &objects.PessoaResult{
							Pessoa:    &p,
							SheetId:   &cfg.id,
							Timestamp: time.Now(),
						})
					}
				}
			case cfg.id + "IFRS- Canoas!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}

					firstCell := row[0].(string)
					if strings.Contains(firstCell, "PESSOAS QUE SAIRAM") {
						break
					}

					nome := strings.Split(firstCell, ". ")
					if len(nome) < 2 {
						continue
					}

					p := objects.Pessoa{
						Abrigo: "Instituto Federal (IFRS) - Canoas",
						Nome:   nome[1],
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Igreja Redenção Nazario!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Igreja Redenção Nazário",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "MODULAR!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Modular",
						Nome:   row[0].(string),
					}
					if len(row) > 2 {
						p.Idade = row[1].(string)
					} else {
						p.Idade = ""
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Paroquia NSRosário!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Paróquia Nossa Senhora do Rosário",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "pediatria HU!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 2 || len(row) < 1 {
						continue
					}
					nome := row[0].(string)
					split := strings.Split(nome, ", ")

					p := objects.Pessoa{
						Abrigo: "Pediatria - Hospital Universitário Canoas",
						Nome:   split[0],
						Idade:  strings.Trim(split[1], ","),
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Rua Itu, 672!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Rua Itu, 672",
						Nome:   strings.TrimRight(row[0].(string), "-"),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "ULBRA!A1:ZZ":
				seen := make(map[string]bool)
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 3 {
						continue
					}

					var p objects.Pessoa
					if len(row) > 4 {
						p = objects.Pessoa{
							Abrigo: removeExtraSpaces("Ulbra" + " " + removeSubstringInsensitive(row[4].(string), "ulbra")),
							Nome:   row[2].(string),
							Idade:  "",
						}
					} else {
						p = objects.Pessoa{
							Abrigo: "Ulbra",
							Nome:   row[2].(string),
							Idade:  "",
						}
					}
					if _, ok := seen[p.Nome]; !ok {
						seen[p.Nome] = true
						serializedData = append(serializedData, &objects.PessoaResult{
							Pessoa:    &p,
							SheetId:   &cfg.id,
							Timestamp: time.Now(),
						})
						fmt.Fprintln(os.Stdout, p)
					}
				}
			case cfg.id + "Unilasalle!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Unilasalle",
						Nome:   row[1].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1-1q4c8Ns6M9noCEhQqBE6gy3FWUv-VQgeUO9c7szGIM" + "SESI!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "SESI",
						Nome:   row[1].(string),
					}
					if len(row) > 2 {
						p.Idade = row[2].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "PARÓQUIA SAO LUIS!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Paróquia São Luis",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1Gf78W5yY0Yiljg-E0rYqbRjxYmBPcG2BtfpGwFk-K5M" + "Página1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 2 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: row[1].(string),
						Nome:   row[0].(string),
						Idade:  "",
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "ENCONTRADOS!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "ULBRA - CANOAS PRÉDIO 11",
						Nome:   row[0].(string),
						Idade:  row[2].(string),
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "CIEP!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 4 || len(row) < 3 {
						continue
					}
					nome := row[2].(string)
					if strings.Contains(nome, "MENOR DE 1") {
						break
					}
					p := objects.Pessoa{
						Abrigo: "CIEP",
						Nome:   nome,
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1RGRoIzSFQaaJF1xZsJhQsMJxXnXWzfZfas29T_PefmY" + "SESI!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 4 || len(row) < 3 {
						continue
					}
					nome := row[2].(string)

					if strings.Contains(nome, "MENOR DE 1") {
						break
					}

					p := objects.Pessoa{
						Abrigo: "SESI",
						Nome:   nome,
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "LIBERATO!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 4 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Liberato",
						Nome:   row[1].(string),
					}

					if len(row) > 5 {
						p.Idade = row[5].(string)
					} else {
						p.Idade = ""
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "SINODAL!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 4 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Sinodal",
						Nome:   row[2].(string),
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}
					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "PARQUE DO TRABALHADOR!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 4 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Parque do Trabalhador",
						Nome:   row[2].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "FENAC II!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 2 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "FENAC",
						Nome:   row[1].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "GINÁSIO DA BRIGADA!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 2 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Ginásio da Brigada, Novo Hamburgo",
						Nome:   row[2].(string),
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "IGREJA NOSSA SENHORA DAS GRAÇAS DA RONDÔNIA!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 2 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Igreja Nossa Senhora das Graças da Rondônia",
						Nome:   row[2].(string),
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "COMUNIDADE SANTO ANTONIO!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 2 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Igreja Santo Antônio - Bairro Liberdade",
						Nome:   row[2].(string),
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "PIO XII!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 2 || len(row) < 3 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Pio XII",
						Nome:   row[2].(string),
					}

					if len(row) > 3 {
						p.Idade = row[3].(string)
					} else {
						p.Idade = ""
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "NOME/ABRIGO!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: row[1].(string) + " Eldorado do Sul",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "05/05 PONTAL!A1:ZZ", cfg.id + "06/05 PONTAL!A1:ZZ":
				for i, row := range content.([][]interface{}) {

					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Pontal do Estaleiro",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "05/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ", cfg.id + "06/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ", cfg.id + "04/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 2 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Gasômetro",
						Nome:   row[1].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}

			case "1yuzazWMydzJKUoBnElV1YTxSKLJsT4fSVHfyJBjLlAY" + "Página1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "SESC Protásio",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Abrigados Lajeado!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 3 {
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
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1O4NqkxHvFDoziS_zClwIjGIAVAGbYkfHTRrM6ogySTo" + "Página1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 3 || len(row) < 1 {
						continue
					}
					p := objects.Pessoa{
						Abrigo: "Venâncio Aires",
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Resgatados!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 5 {
						continue
					}
					var abrigo string
					abrigo = row[4].(string)
					if abrigo == "" {
						abrigo = "Cruzeiro do Sul"
					}

					p := objects.Pessoa{
						Abrigo: abrigo,
						Nome:   row[0].(string),
						Idade:  "",
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1AaQLs2Dqc6lrYstyF8UGLrihCzRRLsy8rlIRixJQ7VU" + "Página1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 3 || len(row) < 1 {
						continue
					}
					var p objects.Pessoa
					pattern := `[0-9]+`
					re := regexp.MustCompile(pattern)
					replacedStr := re.ReplaceAllString(row[0].(string), "")
					if len(replacedStr) > 0 {
						p = objects.Pessoa{
							Abrigo: "Linha Herval - Venâncio Aires",
							Nome:   replacedStr,
							Idade:  "",
						}
					}

					fmt.Fprintln(os.Stdout, p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1IVtSmKRFynQH9I9Cox93YxZe0uwKfjx_CYFzKE96its" + "Sheet1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}
					var p objects.Pessoa
					pattern := `[0-9]+`
					re := regexp.MustCompile(pattern)
					replacedStr := re.ReplaceAllString(row[0].(string), "")
					if len(replacedStr) > 0 {
						p = objects.Pessoa{
							Nome:   replacedStr,
							Idade:  "",
						}
						if row[1].(string) != "" {
							p.Abrigo = row[1].(string)
						} else {
							p.Abrigo = "Abrigo Coelhão"
						}
					}

					fmt.Fprintf(os.Stdout, "%+v\n", p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "16X-68-x7My4u0WEfscL7t4YYw_Ebeco6gaLhE80Q8Wc" + "Página1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 6 {
						continue
					}
					var p objects.Pessoa
					var abrigo string

					abrigo = row[5].(string)
					if abrigo == "" {
						abrigo = "Desconhecido"
					}
					p = objects.Pessoa{
						Abrigo: abrigo,
						Nome:   row[2].(string),
						Idade:  "",
					}


					fmt.Fprintf(os.Stdout, "%+v\n", p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "1wvtgK7ZO9KuJsFDI9syyPWmEyqYoKw2PKssmgfo_jCU" + "Form Responses 1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 4 {
						continue
					}
					var p objects.Pessoa
					var abrigo string
					var nome string

					nome = row[1].(string)
					reg, err := regexp.Compile("[^a-zA-Z\\s]+")
					if err != nil {
						log.Fatal(err)
					}
					nome = reg.ReplaceAllString(nome, "")

					abrigo = row[3].(string)
					if abrigo == "" {
						abrigo = "Desconhecido"
					}
					p = objects.Pessoa{
						Abrigo: abrigo,
						Nome:   nome,
						Idade:  "",
					}


					fmt.Fprintf(os.Stdout, "%+v\n", p)
					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &cfg.id,
						Timestamp: time.Now(),
					})
				}
			case cfg.id + "Sheet1!A1:ZZ":
				for i, row := range content.([][]interface{}) {
					if i < 1 || len(row) < 1 {
						continue
					}

					p := objects.Pessoa{
						Abrigo: row[1].(string),
						Nome:   row[0].(string),
						Idade:  "",
					}

					sheetId := row[2].(string)
					var url string
					if len(row) > 3 {
						url = row[3].(string)
					}

					fmt.Fprintf(os.Stdout, "%+v\n", p)

					serializedData = append(serializedData, &objects.PessoaResult{
						Pessoa:    &p,
						SheetId:   &sheetId,
						URL:       &url,
						Timestamp: time.Now(),
					})
				}
			}
			var cleanedData []*objects.PessoaResult
			for _, pessoa := range serializedData {
				if pessoa.Nome == "" || pessoa.Abrigo == "" || len(strings.Split(pessoa.Nome, " ")) == 1 {
					continue
				}
				cleanedData = append(cleanedData, pessoa)
			}

			if !isDryRun {
				repository.AddToFirestore(cleanedData)
			}

			fmt.Fprintf(os.Stdout, "Scraped data from sheetId %s, range %s. %d results. %d results after cleanup. Dry run? %v", cfg.id, sheetRange, len(serializedData), len(cleanedData), isDryRun)
			// Clearing arrays for next iteration, I don't think this is strictly needed but just in case.
			serializedData = serializedData[:0]
			cleanedData = cleanedData[:0]
			fmt.Fprintln(os.Stdout, "")
		}
	}
}
