package main

import (
	stp "speedup/wordprocess/stopwords"
	"speedup/wordprocess/stringprocess"
	wd "speedup/wordprocess/wordmap"
	"strings"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func main() {

	//fmt.Printf("%s length is %d \n", normStr1, len(str1))

	texto := "Thiago ama tatiane, Tatiane ama thiago. Tatiane e thiago ama isabella"

	stopwords := new(stp.StopWords).InitStopWords(stp.PORTUGUESES)
	wordmap := new(wd.WordMap).InitWordMap()

	words := strings.Split(texto, " ")

	for _, word := range words {

		newWord := stringprocess.ProcessWord(word)

		println(newWord)

		if !stopwords.IsStopWord(newWord) {
			id := wordmap.AddWord(newWord)
			println(id)
		} else {
			println(newWord)
		}
	}

}
