package objects

import (
	"refugio/utils"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (p *PessoaResult) Clean() *PessoaResult {
	p.Nome = cleanNome(p.Nome)
	p.Abrigo = cleanCommon(p.Abrigo)

	return p
}

func (p *PessoaResult) Validate() (bool, *PessoaResult) {
	if p.Nome == "" {
		return false, p
	}
	if p.Abrigo == "" {
		return false, p
	}

	return true, p
}

func cleanNome(name string) string {
	caser := cases.Title(language.BrazilianPortuguese)

	// Basic cleaning
	name = cleanCommon(name)

	// Enforce title case
	name = caser.String(name)

	// Replace long numbers with an empty string. This is to remove sensitive information like phones and document numbers
	regexPhoneNumbers := regexp.MustCompile(`\d{3,}`)
	name = regexPhoneNumbers.ReplaceAllString(name, "")

	return name
}

func cleanCommon(str string) string {
	// Strip leading and trailing whitespace
	str = strings.TrimSpace(str)

	// Remove extra spaces
	str = utils.RemoveExtraSpaces(str)

	// Remove linebreaks, etc
	regexLineBreaks := regexp.MustCompile(`[\r\n\t]+`)
	str = regexLineBreaks.ReplaceAllString(str, "")

	return str
}
