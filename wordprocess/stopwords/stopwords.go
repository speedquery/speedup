package stopwords

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const PORTUGUESES = "portuguese"

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func normalizeText(value string) string {

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normStr1, _, _ := transform.String(t, value)

	return strings.ToLower(normStr1)
}

type StopWords struct {
	stopwords []string
}

func (stp *StopWords) InitStopWords(typ string) *StopWords {

	stp.stopwords = make([]string, 1)

	if typ == PORTUGUESES {

		f, _ := os.Open("speedup/portugues.txt")
		scanner := bufio.NewScanner(f)

		// Set the Split method to ScanWords.
		scanner.Split(bufio.ScanWords)

		// Scan all words from the file.
		for scanner.Scan() {
			line := scanner.Text()
			stp.stopwords = append(stp.stopwords, normalizeText(strings.TrimSpace(line)))
		}
	}

	return stp

}

func (stp *StopWords) IsStopWord(word string) bool {

	for _, value := range stp.GetWords() {
		if word == value {
			return true
		}
	}

	return false
}

func (stp *StopWords) GetWords() []string {

	return stp.stopwords

}
