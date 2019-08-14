package indexwriter

import (
	"encoding/json"
	"fmt"
	doc "speedup/document"
	fs "speedup/filesystem"
	"speedup/wordprocess/stringprocess"
	"strings"
)

type IndexWriter struct {
	//stopwords *stp.StopWords
	fileSystem *fs.FileSystem
}

func (idx *IndexWriter) CreateIndex(fileSystem *fs.FileSystem) *IndexWriter {
	idx.fileSystem = fileSystem
	return idx
}

func (idx *IndexWriter) IndexDocument(document *doc.Document) {

	for attribute, value := range document.GetMap() {

		idAttribute := idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		//println(idatt)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := make([]uint, 0)

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)

			idword := idx.fileSystem.GetWordMap().AddWord(newWord)
			idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)

			wordGroup = append(wordGroup, idword)

			//idx.wordmap.AddWord(newWord)

			//println(newWord)

			//if !stopwords.IsStopWord(newWord) {
			//
			//}
		}

		bolB, _ := json.Marshal(wordGroup)
		fmt.Println(string(bolB))
		idx.fileSystem.GetWordGroupMap().AddAWordGroup(wordGroup)

	}

	//println(idx.fileSystem.GetAttributeWord().ToJson())
	println(idx.fileSystem.GetWordGroupMap().ToJson())

	

	document = nil
}
