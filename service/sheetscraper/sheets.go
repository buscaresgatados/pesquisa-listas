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
			case "ATUALIZADO 06/05!A1:ZZ":
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
			case "ESCOLA ANDRÉ PUENTE!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "EMEF WALTER PERACCHI DE BARCELLOS!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CACHOEIRINHA!A1:ZZ":
				for _, row := range content.([][]interface{}) {
					p := objects.Pessoa{
						Abrigo: row[1].(string),
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
			case "COLÉGIO MARIA AUXILIADORA!A1:ZZ":
				for _, row := range content.([][]interface{}) {
					p := objects.Pessoa{
						Abrigo: "Colégio Maria Auxiliadora",
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
			case "ULBRA - Prédio 14!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "COLÉGIO MIGUEL LAMPERT!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "AMORJI!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "ESCOLA RONDONIA!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Escola Jacob Longoni!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "COLÉGIO ESPÍRITO SANTO!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CLUBE DOS EMPREGADOS DA PETROBRÁS!A1:ZZ":
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
							SheetId:   cfg.id,
							Timestamp: time.Now(),
						})
					}
				}
			case "Colegio Guajuviras!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CEL São José!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CR BRASIL!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CSSGAPA!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CTG Brazão do Rio Grande!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "CTG Seiva Nativa!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "EMEF ILDO!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Escola Irmao pedro!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "FENIX!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "PARÓQUIA SANTA LUZIA!A1:ZZ":
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
							SheetId:   cfg.id,
							Timestamp: time.Now(),
						})
					}
				}
			case "IFRS- Canoas!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Igreja Redenção Nazario!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "MODULAR!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Paroquia NSRosário!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "pediatria HU!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Rua Itu, 672!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "ULBRA!A1:ZZ":
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
							SheetId:   cfg.id,
							Timestamp: time.Now(),
						})
						fmt.Fprintln(os.Stdout, p)
					}
				}
			case "Unilasalle!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "SESI!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "PARÓQUIA SAO LUIS!A1:ZZ":
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
						SheetId:   cfg.id,
						Timestamp: time.Now(),
					})
				}
			case "Página1!A1:ZZ":
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
						SheetId:   cfg.id,
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
