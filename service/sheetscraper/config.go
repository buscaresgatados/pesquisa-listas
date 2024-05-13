package sheetscraper

type SheetConfig struct {
	id          string
	sheetRanges []string
	name        string
}

var Config []SheetConfig = []SheetConfig{
	// SEM ACESSO -- ENTRAR EM CONTATO COM O PROPRIETÁRIO
	// {
	// 	id:          "1Kw8_Tl4cE4_hrb2APfSlNRli7IxgBbwGXq9d7aNSTzE",
	// 	sheetRanges: []string{"Cadastro inicial!A1:ZZ"},
	// 	name:        "Cadastro de Abrigados Escola Aurélio Reis",
	// },
	{
		id:          "1--z2fbczdFT4RSoji7jXc2jDDU5HqWgAU93NuROBQ78",
		sheetRanges: []string{"Queila!A1:ZZ"},
		name:        "GRAVATAÍ - Lista de acolhidos por abrigo",
	},
	{
		id:          "14WIowAKQo5o_FviBw_6hRxnzAclw5xTvHbUiQuU8qDw",
		sheetRanges: []string{"Cadastro!A1:ZZ"},
		name:        "Alojados Elyseu",
	// },
	{
		id:          "1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA",
		sheetRanges: []string{"Alojados!A1:ZZ"},
		name:        "Lista de alojados - To Salvo Vale dos Sinos",
	},
	{
		id:          "1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ",
		sheetRanges: []string{"CADASTRO_ABRIGADOS!A1:ZZ"},
		name:        "LISTA DESABRIGADOS",
	},
	{
		id:          "1zt_yrzvU2nmihyZG7rR67iqlkwlcDz3LjSb1UBynQ3c",
		sheetRanges: []string{"ALOJADOS x ABRIGOS!A1:ZZ"},
		name:        "SAO LEOPOLDO - LISTA ALOJADOS",
	},
	{
		id:          "1frgtJ9eK05OqsyLwOBiZ2Q6E7e4_pWyrb7fJioqfEMs",
		sheetRanges: []string{"ATUALIZADO 06/05!A1:ZZ"},
		name:        "RESGATADOS - Enchente RS",
	},
	{
		id: "1-1q4c8Ns6M9noCEhQqBE6gy3FWUv-VQgeUO9c7szGIM",
		sheetRanges: []string{
	"COLÉGIO ADVENTISTA DE CANOAS - CACN!A1:ZZ",
	"ESCOLA ANDRÉ PUENTE!A1:ZZ",
	"EMEF WALTER PERACCHI DE BARCELLOS!A1:ZZ",
	"CACHOEIRINHA!A1:ZZ",
	"ULBRA - Prédio 14!A1:ZZ",
	"COLÉGIO MIGUEL LAMPERT!A1:ZZ",
	"AMORJI!A1:ZZ",
	"ESCOLA RONDONIA!A1:ZZ",
	"Escola Jacob Longoni!A1:ZZ",
	"COLÉGIO ESPÍRITO SANTO!A1:ZZ",
	"CLUBE DOS EMPREGADOS DA PETROBRÁS!A1:ZZ",
	"Colegio Guajuviras!A1:ZZ",
	"CEL São José!A1:ZZ",
	"CR BRASIL!A1:ZZ",
	"CSSGAPA!A1:ZZ",
	"CTG Brazão do Rio Grande!A1:ZZ",
	"CTG Seiva Nativa!A1:ZZ",
	"EMEF ILDO!A1:ZZ",
	"Escola Irmao pedro!A1:ZZ",
	"FENIX!A1:ZZ",
	"PARÓQUIA SANTA LUZIA!A1:ZZ",
	"IFRS- Canoas!A1:ZZ",
	"Igreja Redenção Nazario!A1:ZZ",
	"MODULAR!A1:ZZ",
	"Paroquia NSRosário!A1:ZZ",
	"pediatria HU!A1:ZZ",
			"Rua Itu, 672!A1:ZZ",
			"ULBRA!A1:ZZ",
			"Unilasalle!A1:ZZ",
			"SESI!A1:ZZ",
			"PARÓQUIA SAO LUIS!A1:ZZ",
		},
		name: "ABRIGADOS EM CANOAS 01",
	},
	{
		id:          "1Gf78W5yY0Yiljg-E0rYqbRjxYmBPcG2BtfpGwFk-K5M",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "LISTA GERAL EM ATUALIZAÇÃO CONSTANTE",
	},
	{
		id:          "1T_yd-M6BG1qYdQKeMo2U_AffqRCxkExqpB39iQXig5s",
		sheetRanges: []string{"ENCONTRADOS!A1:ZZ"},
		name:        "ULBRA - CANOAS PRÉDIO 11",
	},
	{
		id: "1RGRoIzSFQaaJF1xZsJhQsMJxXnXWzfZfas29T_PefmY",
		sheetRanges: []string{
	"CIEP!A1:ZZ",
	"LIBERATO!A1:ZZ",
	"SINODAL!A1:ZZ",
	"PARQUE DO TRABALHADOR!A1:ZZ",
	"SESI!A1:ZZ",
	"FENAC II!A1:ZZ",
	"GINÁSIO DA BRIGADA!A1:ZZ",
	"IGREJA NOSSA SENHORA DAS GRAÇAS DA RONDÔNIA!A1:ZZ",
	"COMUNIDADE SÃO FRANCISCO!A1:ZZ",
	"COMUNIDADE SANTO ANTONIO!A1:ZZ",
	"PIO XII!A1:ZZ",
		"LISTA MULHERES!A1:ZZ",
		"IGREJA NOSSA SENHORA DAS GRAÇAS !A1:ZZ",
	},
	name: "ACOLHIDOS NOS ALOJAMENTOS DA PMNH",
	},
	{
		id:          "1hKJVs-RLiSUpx-1Rd9wS1k8RLqxkPWK4hNob-t8v2Ko",
		sheetRanges: []string{"NOME/ABRIGO!A1:ZZ"},
		name:        "LISTA RESGATADOS ELDORADO - 04/05",
	},
	{id: "1-xhmS1VQ95LFI05XG8o9JO3mPk8KxQtrxAZe4GNYO3I",
		sheetRanges: []string{
			"05/05 PONTAL!A1:ZZ",
			"05/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
			"06/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
			"04/05 GASÔMETRO (NÃO MEXER!)!A1:ZZ",
		},
		name: "RESGATADOS PONTAL",
	},
	{
		id:          "1yuzazWMydzJKUoBnElV1YTxSKLJsT4fSVHfyJBjLlAY",
		sheetRanges: []string{"Lista Abrigados!A1:ZZ"},
		name:        "Lista desabrigados Sesc Protásio",
	},
	{
		id:          "1bNw-t0RUE-AP-2quCU80w5meGYVtD7WjjuESb7tXxTo",
		sheetRanges: []string{"Abrigados Lajeado!A1:ZZ"},
		name:        "Abrigados Lajeado",
	},
	{
		id:          "1O4NqkxHvFDoziS_zClwIjGIAVAGbYkfHTRrM6ogySTo",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Abrigados - Venâncio Aires",
	},
	{
		id:          "1Pd8NVuEtnR7-IlLF7cJ3XY7yVSoJMgY47-eepe2BBXo",
		sheetRanges: []string{"Resgatados!A1:ZZ"},
		name:        "RESGATADOS - Cruzeiro do Sul",
	},
	{
		id:          "1AaQLs2Dqc6lrYstyF8UGLrihCzRRLsy8rlIRixJQ7VU",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Abrigados - Linha Herval",
	},
	{
		id:          "1ym1_GhBA47LhH97HhggICESiUbKSH-e2Oii1peh6QF0",
		sheetRanges: []string{"Sheet1!A1:ZZ"},
		name:        "Compilado XLSX",
	},
	{
		id:          "1IVtSmKRFynQH9I9Cox93YxZe0uwKfjx_CYFzKE96its",
		sheetRanges: []string{"Sheet1!A1:ZZ"},
		name:        "ALOJAMENTO COELHÃO",
	},
	{
		id:          "16X-68-x7My4u0WEfscL7t4YYw_Ebeco6gaLhE80Q8Wc",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Resgatados viaduto da Dom Pedro",
	},
	{
		id:          "1wvtgK7ZO9KuJsFDI9syyPWmEyqYoKw2PKssmgfo_jCU",
		sheetRanges: []string{"Form Responses 1!A1:ZZ"},
		name:        "localizacao desabrigados canoas (Responses)-elaborada por voluntarios da prefeitura de Canoas",
	},
	{
		id:          "1fH7OA5bnY5OLfY7Xis6bVQq12VIhS_VIyYYekPBr5NA",
		sheetRanges: []string{"Respostas ao formulário 1!A1:ZZ"},
		name:        "Lista Abrigados em Cerro Grande do Sul",
	},
	{
		id:          "1T_yd-M6BG1qYdQKeMo2U_AffqRCxkExqpB39iQXig5s",
		sheetRanges: []string{"ENCONTRADOS!A1:ZZ"},
		name:        "ULBRA - CANOAS PRÉDIO 11",
	},
	{
		id: "1eC6z6RPNNarLMSqVqU-FQOHopCKWCN4CFDn34uTYGcA",
		sheetRanges: []string{
			"Página 1!A1:ZZ",
			"Página2!A1:ZZ",
			"Página 3!A1:ZZ",
			"Página 4!A1:ZZ",
			"Página 5!A1:ZZ",
		},
		name: "PESSOAS ALOJADAS EM SENTINELA DO SUL",
	},
	{
		id:          "1LdM2ZvYBNdtKekLgHPRs6lg9VGpD-7wBSZsE5c5Mptk",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Abrigados Estrela",
	},
	{
		id:          "1-cA0MB_1aQTOtXVL2pyPWSXjuTMg6U1PsyBAICjdGxo",
		sheetRanges: []string{"Gravataí!A1:ZZ"},
		name:        "GRAVATAÍ PESSOAS RESGATADAS",
	},
	{
		id:          "16rN5pniNiIsbJAv25A0AfW5SdccJjPVDov7EDqwDOQM",
		sheetRanges: []string{"Abrigados!A1:ZZ"},
		name:        "Resgates - Abrigo Julio de Castílhos",
	},
	{
		id:          "1gfQ28EPN99LQaZqZzMeB-pdxgK9SST1OYy-jTOl7rdk",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "PAROQUIA NOSSA SENHORA APARECIDA",
	},
	{
		id:          "1KgPjNIDQOmDA59A8u4HIOzsL41ZGQH97n-2jl99tfuU",
		sheetRanges: []string{"Sheet1!A1:ZZ"},
		name:        "Lista Abrigo Liberato",
	},
	{
		id:          "1VE_WnX5MuVF_4Mtos7a-S7eYPrudeygv-OWddwuCkYc",
		sheetRanges: []string{"Giovana!A1:ZZ"},
		name:        "DIGITALIZAÇÃO DE REGISTRO DE RESGATES - RS",
	},
	{
		id:          "1xaEPlk8JonATIOAvQEc0Dev-QVAzx2AwUzLHBhbA3rI",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Resgates no viaduto da Santa Rita -> Eldorado - Guaiba",
	},
	{
		id: "1FRHLIpLOE0xr7IwecZHU6Q6QMkescPuqjtxmjIb2GI8",
		sheetRanges: []string{
			"AD55!A1:ZZ",
			"CESE!A1:ZZ",
			"Comunidade Santa Clara!A1:ZZ",
			"CTG Guapos da Amizade!A1:ZZ",
			"Gaditas!A1:ZZ",
			"Ginásio Placar!A1:ZZ",
			"ONG Vida Viva!A1:ZZ",
			"Onze Unidos!A1:ZZ",
			"CTG Carreteiros!A1:ZZ",
			"Abrigo Santa Clara!A1:ZZ",
			"SESI!A1:ZZ",
			"Paróquia Santa Luzia (bairro Fátima)!A1:ZZ",
			"Igreja Betel!A1:ZZ",
		},
		name: "Abrigos - Cachoeirinha/RS",
	},
	{
		id:          "1TVv1WEjrPBpnKsFIV60jz0kWPK6idovmnJDaGg6KKXw",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Abrigados Porto Novo/SESI",
	},
	{
		id: "1kKfTi8N-XL2bcML8Xtf3cT1FNIzinqh4woHDjHn2Bgs",
		sheetRanges: []string{
			"ATUALIZADO 05/05!A1:ZZ",
			"ATUALIZADO 06/05!A1:ZZ",
			"ATUALIZADO 07/05!A1:ZZ",
			"ATUALIZADO 08/05!A1:ZZ",
		},
		name: "NOVA de RESGATADOS - RS",
	},
	{
		id: "1K3DRVlSpK3tWQ1B83Q9pxkhSivIsmf38FTb6SVjMzT4",
		sheetRanges: []string{
			"Resgatados Prefeitura SL!A1:ZZ",
			"RESGATADOS/ABRIGADOS!A1:ZZ",
			"Resgatados - Fernanda!A1:ZZ",
		},
		name: "Respostas Pedido Desaparecidos Enchente",
	},
	{
		id: "1TvBXpT1vZpuAffc2rb8VE2mBMEFnG1_sqIlIL4b1PuA",
		sheetRanges: []string{
			"Velha Cambona!A1:ZZ",
			"NSra Fátima!A1:ZZ",
			"Vila Rica!A1:ZZ",
		},
		name: "Abrigos Portão - RS",
	},
	{
		id: "1q3Z2iX_vop9EumvB-4UyZsVQl58ZQ0M1JnwQsc6HAAo",
		sheetRanges: []string{
			"06/05!A1:ZZ",
			"07/05!A1:ZZ",
			"08/05!A1:ZZ",
		},
		name: "[atualizada 08/05/2024 às 16h25] RESGATADOS - Bairro Humaitá Porto Alegre",
	},
	{
		id:          "17GlFds1C-sdRdpWkZczzisTdItbdWgVAMXwXV60htyA",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Abrigados CESMAR",
	},
	{
		id:          "10OnXFy-8TtUr3gw9yvtWroI7Z1psXGjdyBA3KMQKstE",
		sheetRanges: []string{"Planilha1!A1:ZZ"},
		name:        "Abrigados - FAPA",
	},
	{
		id:          "1oMPwqFsfjlHB1snApt_BGGJrwTSmFn_R8_4Bm7ufAoY",
		sheetRanges: []string{"Página1!A1:ZZ"},
		name:        "Desabrigados - WhatsApp Bot",
	},
}
