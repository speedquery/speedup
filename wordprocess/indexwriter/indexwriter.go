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

		idAttribute := idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		//println(*idAttribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := make([]*uint, 0)

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)
			idword := idx.fileSystem.GetWordMap().AddWord(newWord)
			idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)

			wordGroup = append(wordGroup, idword)

		}

		//bolB, _ := json.Marshal(wordGroup)
		//fmt.Println(string(bolB))
		idWordGroup := idx.fileSystem.GetWordGroupMap().AddAWordGroup(wordGroup)
		idx.fileSystem.GetAttributeGroupWord().AddGroupWordsOfAttribute(idAttribute, idWordGroup)

		idDocument := document.GetID()
		idx.fileSystem.GetGroupWordDocument().AddGroupWordDocument(&idDocument, idWordGroup)

	}

	//println(idx.fileSystem.GetAttributeWord().ToJson())
	println("ATT MAP", idx.fileSystem.GetAttributeMap().ToJson())
	println("WORD MAP", idx.fileSystem.GetWordMap().ToJson())

	println(idx.fileSystem.GetWordGroupMap().ToJson())
	println(idx.fileSystem.GetGroupWordDocument().ToJson())

}
