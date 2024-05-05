package sheetscraper

type SheetConfig struct {
	id          string
	sheetRanges []string
	name        string
}

var Config []SheetConfig = []SheetConfig{
	{
		id:          "1hxzHYE4UR1YbcH3ZQoPfcQTgPbRm5T6lShkDwGDoeXA",
		sheetRanges: []string{"Estância Velha!A1:ZZ"},
		name:        "Lista de alojados - To Salvo Vale dos Sinos",
	},
}
