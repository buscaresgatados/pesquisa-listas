package sheetscraper

type SheetConfig struct {
	id          string
	sheetRanges []string
	name        string
}

var Config []SheetConfig = []SheetConfig{
	{
		id:          "1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA",
		sheetRanges: []string{"Alojados!A1:ZZ"},
		name:        "Lista de alojados - To Salvo Vale dos Sinos",
	},
	{
		id:          "1f5gofOOv4EFYWhVqwPWbgF2M-7uHrJrCMiP7Ug4y6lQ",
		sheetRanges: []string{"CADASTRO_EM_TEMPO_REAL!A1:ZZ"},
		name:        "LISTA DESABRIGADOS",
	},
	{
		id:          "1zt_yrzvU2nmihyZG7rR67iqlkwlcDz3LjSb1UBynQ3c",
		sheetRanges: []string{"ALOJADOS x ABRIGOS!A1:ZZ"},
		name:        "SAO LEOPOLDO - LISTA ALOJADOS",
	},
	{
		id:          "1frgtJ9eK05OqsyLwOBiZ2Q6E7e4_pWyrb7fJioqfEMs",
		sheetRanges: []string{"PESSOAS RESGATADAS"},
		name:        "RESGATADOS - Enchente RS",
	},
}
