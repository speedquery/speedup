package indexwriter

import (
	"container/list"
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

func (idx *IndexWriter) IndexDocument(document *doc.Document, onExit func()) {

	defer onExit()
	//println("Documento", document.GetID())

	for attribute, value := range document.GetMap() {

		idAttribute := idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		//println(*idAttribute, attribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := list.New()

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)
			idword := idx.fileSystem.GetWordMap().AddWord(newWord)
			idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)
			wordGroup.PushBack(idword)

		}

		//bolB, _ := json.Marshal(wordGroup)
		//idx.fileSystem.GetWordGroupMap().AddAWordGroup(wordGroup)
		//fmt.Println(*idWordGroup, string(bolB))
		//idx.fileSystem.GetAttributeGroupWord().AddGroupWordsOfAttribute(idAttribute, idWordGroup)

		//idDocument := document.GetID()
		//println("DOCUMENTO GRUPO", idDocument, *idWordGroup)
		//idx.fileSystem.GetGroupWordDocument().AddGroupWordDocument(idWordGroup, idDocument)
		//println(*idWordGroup, idDocument)

	}

	//println(idx.fileSystem.GetAttributeWord().ToJson())
	//println("ATT MAP", idx.fileSystem.GetAttributeMap().ToJson())
	//println("WORD MAP", idx.fileSystem.GetWordMap().ToJson())
	//println("ATT WORD", idx.fileSystem.GetAttributeWord().ToJson())
	//println("WORD GROUP MAP", idx.fileSystem.GetWordGroupMap().ToJson())
	//println("DOCUMENT GROUP", idx.fileSystem.GetGroupWordDocument().ToJson())

	//println(idx.fileSystem.GetWordGroupMap().ToJson())
	//println(idx.fileSystem.GetGroupWordDocument().ToJson())

}
