package sheetscraper

type SheetConfig struct {
	id          string
	sheetRanges []string
	name        string
}

var Config []SheetConfig = []SheetConfig{
	// {
	// 	id:          "1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA",
	// 	sheetRanges: []string{"Alojados!A1:ZZ"},
	// 	name:        "Lista de alojados - To Salvo Vale dos Sinos",
	// },
	// {
	// 	id:          "1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ",
	// 	sheetRanges: []string{"CADASTRO_EM_TEMPO_REAL!A1:ZZ"},
	// 	name:        "LISTA DESABRIGADOS",
	// },
	// {
	// 	id:          "1zt_yrzvU2nmihyZG7rR67iqlkwlcDz3LjSb1UBynQ3c",
	// 	sheetRanges: []string{"ALOJADOS x ABRIGOS!A1:ZZ"},
	// 	name:        "SAO LEOPOLDO - LISTA ALOJADOS",
	// },
	// {
	// 	id:          "1frgtJ9eK05OqsyLwOBiZ2Q6E7e4_pWyrb7fJioqfEMs",
	// 	sheetRanges: []string{"ATUALIZADO 06/05!A1:ZZ"},
	// 	name:        "RESGATADOS - Enchente RS",
	// },
	// {
	// 	id: "1-1q4c8Ns6M9noCEhQqBE6gy3FWUv-VQgeUO9c7szGIM",
	// 	sheetRanges: []string{
	// 		"ESCOLA ANDRÉ PUENTE!A1:ZZ",
	// 		"EMEF WALTER PERACCHI DE BARCELLOS!A1:ZZ",
	// 		"CACHOEIRINHA!A1:ZZ",
	// 		"ULBRA - Prédio 14!A1:ZZ",
	// 		"COLÉGIO MIGUEL LAMPERT!A1:ZZ",
	// 		"AMORJI!A1:ZZ",
	// 		"ESCOLA RONDONIA!A1:ZZ",
	// 		"Escola Jacob Longoni!A1:ZZ",
	// 		"COLÉGIO ESPÍRITO SANTO!A1:ZZ",
	// 		"CLUBE DOS EMPREGADOS DA PETROBRÁS!A1:ZZ",
	// 		"Colegio Guajuviras!A1:ZZ",
	// 		"CEL São José!A1:ZZ",
	// 		"CR BRASIL!A1:ZZ",
	// 		"CSSGAPA!A1:ZZ",
	// 		"CTG Brazão do Rio Grande!A1:ZZ",
	// 		"CTG Seiva Nativa!A1:ZZ",
	// 		"EMEF ILDO!A1:ZZ",
	// 		"Escola Irmao pedro!A1:ZZ",
	// 		"FENIX!A1:ZZ",
	// 		"PARÓQUIA SANTA LUZIA!A1:ZZ",
	// 		"IFRS- Canoas!A1:ZZ",
	// 		"Igreja Redenção Nazario!A1:ZZ",
	// 		"MODULAR!A1:ZZ",
	// 		"Paroquia NSRosário!A1:ZZ",
	// 		"pediatria HU!A1:ZZ",
	// 		"Rua Itu, 672!A1:ZZ",
	// 		"ULBRA!A1:ZZ",
	// 		"Unilasalle!A1:ZZ",
	// 		"SESI!A1:ZZ",
	// 		"PARÓQUIA SAO LUIS!A1:ZZ",
	// 	},
	// 	name: "ABRIGADOS EM CANOAS 01",
	// },
	// {
	// 	id:          "1Gf78W5yY0Yiljg-E0rYqbRjxYmBPcG2BtfpGwFk-K5M",
	// 	sheetRanges: []string{"Página1!A1:ZZ"},
	// 	name:        "LISTA GERAL EM ATUALIZAÇÃO CONSTANTE",
	// },
	// {
	// 	id:          "1T_yd-M6BG1qYdQKeMo2U_AffqRCxkExqpB39iQXig5s",
	// 	sheetRanges: []string{"ENCONTRADOS!A1:ZZ"},
	// 	name:        "ULBRA - CANOAS PRÉDIO 11",
	// },
	// {
	// 	id: "1RGRoIzSFQaaJF1xZsJhQsMJxXnXWzfZfas29T_PefmY",
	// 	sheetRanges: []string{
	// 		"CIEP!A1:ZZ",
	// 		"LIBERATO!A1:ZZ",
	// 		"SINODAL!A1:ZZ",
	// 		"PARQUE DO TRABALHADOR!A1:ZZ",
	// 		"SESI!A1:ZZ",
	// 		"FENAC II!A1:ZZ",
	// 		"GINÁSIO DA BRIGADA!A1:ZZ",
	// 		"IGREJA NOSSA SENHORA DAS GRAÇAS DA RONDÔNIA!A1:ZZ",
	// 		"COMUNIDADE SÃO FRANCISCO!A1:ZZ",
	// 		"COMUNIDADE SANTO ANTONIO!A1:ZZ",
	// 		"PIO XII!A1:ZZ",
	// 	},
	// 	name: "ACOLHIDOS NOS ALOJAMENTOS DA PMNH",
	// },
	// {
	// 	id:          "1hKJVs-RLiSUpx-1Rd9wS1k8RLqxkPWK4hNob-t8v2Ko",
	// 	sheetRanges: []string{"NOME/ABRIGO!A1:ZZ"},
	// 	name:        "LISTA RESGATADOS ELDORADO - 04/05",
	// },
	// {id: "1-xhmS1VQ95LFI05XG8o9JO3mPk8KxQtrxAZe4GNYO3I",
	// 	sheetRanges: []string{
	// 		"05/05 PONTAL!A1:ZZ",
	// 		"06/05 PONTAL!A1:ZZ",
	// 		"05/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
	// 		"06/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
	// 		"04/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
	// 	},
	// 	name: "RESGATADOS PONTAL",
	// },
	// {
	// 	id:          "1yuzazWMydzJKUoBnElV1YTxSKLJsT4fSVHfyJBjLlAY",
	// 	sheetRanges: []string{"Página1!A1:ZZ"},
	// 	name:        "Lista desabrigados Sesc Protásio",
	// },
	// {
	// 	id:          "1bNw-t0RUE-AP-2quCU80w5meGYVtD7WjjuESb7tXxTo",
	// 	sheetRanges: []string{"Abrigados Lajeado!A1:ZZ"},
	// 	name:        "Abrigados Lajeado",
	// },
	// {
	// 	id:          "1O4NqkxHvFDoziS_zClwIjGIAVAGbYkfHTRrM6ogySTo",
	// 	sheetRanges: []string{"Página1!A1:ZZ"},
	// 	name:        "Abrigados - Venâncio Aires",
	// },
	// {
	// 	id:          "1Pd8NVuEtnR7-IlLF7cJ3XY7yVSoJMgY47-eepe2BBXo",
	// 	sheetRanges: []string{"Resgatados!A1:ZZ"},
	// 	name:        "RESGATADOS - Cruzeiro do Sul",
	// },
	// {
	// 	id:          "1AaQLs2Dqc6lrYstyF8UGLrihCzRRLsy8rlIRixJQ7VU",
	// 	sheetRanges: []string{"Página1!A1:ZZ"},
	// 	name:        "Abrigados - Linha Herval",
	// },
	// {
	// 	id:          "1ym1_GhBA47LhH97HhggICESiUbKSH-e2Oii1peh6QF0",
	// 	sheetRanges: []string{"Sheet1!A1:ZZ"},
	// 	name:        "Compilado XLSX",
	// },
	// {
	// 	id: "16X-68-x7My4u0WEfscL7t4YYw_Ebeco6gaLhE80Q8Wc",
	// 	sheetRanges: []string{"Página1!A1:ZZ"},
	// 	name: "Resgatados viaduto da Dom Pedro",
	// },
	// {
	// id:          "1IVtSmKRFynQH9I9Cox93YxZe0uwKfjx_CYFzKE96its",
	// sheetRanges: []string{"Sheet1!A1:ZZ"},
	// name:        "ALOJAMENTO COELHÃO",
	// },
}
