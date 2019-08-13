package indexwriter

import (
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

		idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)

			idx.fileSystem.GetWordMap().AddWord(newWord)

			//idx.wordmap.AddWord(newWord)

			println(newWord)

			//if !stopwords.IsStopWord(newWord) {
			//
			//}
		}

	}

	document = nil
}