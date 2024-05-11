package objects

import "time"

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
	Trace  string `json:"logging.googleapis.com/trace"`
	UserIP string `json:"user_ip"`
}

type Source struct {
	Nome    string
	SheetId string
	URL     string
}
