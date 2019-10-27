package stringprocess

import (
	"log"
	"regexp"
	"speedup/utils"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func NormalizeText(value string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normStr1, _, _ := transform.String(t, value)

	return strings.ToLower(reg.ReplaceAllString(normStr1, ""))
}

func ProcessWord(word string) string {

	word = strings.TrimSpace(word)

	if utils.IsNumber(word) {
		return word
	}

	newWord := make([]string, 5)
	tokens := strings.Split(word, "")

	for _, value := range tokens {
		value = NormalizeText(value)
		newWord = append(newWord, value)
	}

	return strings.Join(newWord, "")

}
