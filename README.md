# Contribuindo com o projeto

## Ferramentas necessárias
- [Golang 1.22.2](https://go.dev/doc/install)
- IDE ou editor de texto (vs-code, goland, etc)

## Passos
- Clonar este repositório: https://github.com/buscaresgatados/pesquisa-listas/
- Acessar este trello: https://trello.com/b/aZRYM4UL/busca-resgatados
- Baixar o arquivo `config.zip` disponível no grupo do discord, no canal #docs
- Descompactar estes arquivos no root da pasta `service`

## Editando arquivos
Os únicos arquivos necessários para o _scraping_ das planilhas são:
- `service/sheetscraper/config.go`
- `service/sheetscraper/sheets.go`

### config.go
Criar um novo objeto ao final do arquivo, com a seguinte estrutura:
```
{
    id: "<ID_DA_PLANILHA>",
    sheetRanges: []string{"<NOME_DA_ABA>!A1:ZZ"},
    name: "<NOME_DA_PLANILHA>",
},
```
> **_NOTE:_**  O ID da planilha está depois de `/d/`:<br>
> https://docs.google.com/spreadsheets/d/abc123456/edit?pli=1#gid=972231790.<br>
> Neste caso: `abc123456`


### sheets.go
Criar um novo `case` dentro da função `Scrape`, pelo `<ID_DA_PLANILHA> + <NOME_DA_ABA!A1:ZZ>`. Exemplo:
```go
case  "1-cA0MB_1aQTOtXVL2pyPWSXjuTMg6U1PsyBAICjdGxo"  +  "Sheet1!A1:ZZ":
	for  i, row  :=  range  content.([][]interface{}) {
		if  i  <  1  ||  len(row) <  3 {
			continue
		}
	p  :=  objects.Pessoa{
		Abrigo: row[2].(string),
		Nome: row[1].(string),
		Idade: "",
	}

	fmt.Fprintln(os.Stdout, p)

	serializedData  =  append(serializedData, &objects.PessoaResult{
		Pessoa: &p,
		SheetId: &cfg.id,
		Timestamp: time.Now(),
	})
}
```

## Running the script
Para testar o script, rode esse comando:<br>
`export $(cat .env | xargs ) && go build -o app && ./app scrape --isDryRun=true`

O output será printado no console. Ele deverá conter a seguinte estrutura:
```
{Abrigo:Associacao Nome:Eva Tavares Idade:}
{Abrigo:Associacao Nome:Cristiano Camargo Idade:}
{Abrigo:Associacao Nome:Nara Regina Pedroso Idade:}
```

Após validar que a estrutura está correta, o script deve ser rodado sem _dryRun_:<br>
`export $(cat .env | xargs ) && go build -o app && ./app scrape --isDryRun=false`

Isso vai fazer com que os dados sejam salvos no Banco de Dados.<br>
Após rodar o script, commitar no repositório as mudanças.<br><br>

Obrigado por contribuir!
