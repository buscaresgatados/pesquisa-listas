package objects

import "time"

type Pessoa struct {
	Abrigo string
	Nome   string
	Idade  string
}

type PessoaResult struct {
	*Pessoa
	SheetId   *string
	URL       *string
	Timestamp time.Time
}

type PessoaReturn struct {
	results []PessoaResult
}

type PessoaSearchResult struct {
	ObjectID string
	Nome     string
}
