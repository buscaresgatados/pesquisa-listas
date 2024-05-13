package objects

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Abrigo     string
	Nome       string
	Idade      string
	Observacao string
}

type PessoaResult struct {
	*Pessoa
	SheetId   *string
	URL       *string
	Timestamp time.Time
}

type PessoaSearchResult struct {
	ObjectID string
	Nome     string
}

type PessoaCountResult struct {
	Total int `json:"total_records"`
}

type PessoaMostRecentResult struct {
	Timestamp *time.Time `json:"timestamp"`
}

type AccessLog struct {
	Trace   *string `json:"logging.googleapis.com/trace"`
	UserIP  *string `json:"user_ip"`
	KeyUser *string `json:"key_user"`
}

type Source struct {
	Nome    string
	SheetId string
	URL     string
	Observacao	string
	Sheets  []string
}

func (s *Source) String() string {
    return fmt.Sprintf("URL: %s, SheetId: %s, Sheets: %v", s.URL, s.SheetId, s.Sheets)
}
