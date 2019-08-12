package main

import (
	"fmt"
	"speedup/document"
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
	doc := new(document.Document).CreateDocument(1)
	doc.AddField("nome", "thiago luiz rodrigues")
	doc.AddField("idade", "32")

	//criar uma função chamada indexDocument
	stopwords := new(stp.StopWords).InitStopWords(stp.PORTUGUESES)
	wordmap := new(wd.WordMap).InitWordMap()

	for _, value := range doc.GetMap() {

		formatedValue := fmt.Sprintf("%v", value)

		words := strings.Split(formatedValue, " ")

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)

			if !stopwords.IsStopWord(newWord) {
				wordmap.AddWord(newWord)
				println(word)
			}
		}

	}

}
