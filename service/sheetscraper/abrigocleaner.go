package sheetscraper

const (
	AbrigoDeduplicationSheetId = "1d15WX-dSXnnwVHNMpO6PxqQtf1CrHZq_T6ymrwSsp8c"
	AbrigoDeduplicationRange   = "Sheet3"
)

func getAbrigosMapping() map[string]string {
	ss := SheetsSource{}
	content, err := ss.Read(AbrigoDeduplicationSheetId, AbrigoDeduplicationRange)
	if err != nil {
		panic(err)
	}
	abrigoDeduplicationMap := make(map[string]string)

	for i, row := range content.([][]interface{}) {
		if i == 0 || len(row) < 2 {
			continue
		}
		abrigoDeduplicationMap[row[0].(string)] = row[1].(string)
	}
	return abrigoDeduplicationMap
}
