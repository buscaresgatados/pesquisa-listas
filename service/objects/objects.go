package objects

import "time"

type Pessoa struct {
	Abrigo string
	Nome   string
	Idade  string
}

type SheetId string

type PessoaResult struct {
	*Pessoa
	SheetId   string
	Timestamp time.Time
}

type PessoaReturn struct {
	results []PessoaResult
}
