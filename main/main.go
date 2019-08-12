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

	doc = new(document.Document).CreateDocument(2)
	doc.AddField("nome", "thiago luiz rodrigues")
	doc.AddField("idade", 32)

	json := doc.ToJson()
	println(json)

	mp := doc.ToMap(`{"_key":1,"idade":32,"nome":"thiago luiz rodrigues"}`)
	fmt.Println(mp)

	texto := "Thiago ama tatiane, Tatiane ama thiago. Tatiane e thiago ama isabella"

	stopwords := new(stp.StopWords).InitStopWords(stp.PORTUGUESES)
	wordmap := new(wd.WordMap).InitWordMap()

	words := strings.Split(texto, " ")

	for _, word := range words {

		newWord := stringprocess.ProcessWord(word)

		if !stopwords.IsStopWord(newWord) {
			wordmap.AddWord(newWord)
			//println(id)
		}
	}

}
