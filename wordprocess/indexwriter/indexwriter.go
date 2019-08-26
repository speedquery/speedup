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

func (idx *IndexWriter) IndexDocument(document *doc.Document, onExit func()) {

	defer onExit()
	//println("Documento", document.GetID())

	for attribute, value := range document.GetMap() {

		idAttribute := idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := make([]string, 0) //list.New()

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)
			idword := idx.fileSystem.GetWordMap().AddWord(newWord)
			idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)
			wordGroup = append(wordGroup, fmt.Sprint(*idword))

		}

		//strings.Join()
		//justString :=

		//fmt.Println(justString)
		/**
		h := sha1.New()
		h.Write([]byte("s"))
		sha1_hash := hex.EncodeToString(h.Sum(nil))
		println(sha1_hash)
		**/
		//bolB, _ := json.Marshal(wordGroup)
		idWordGroup := idx.fileSystem.GetWordGroupMap().AddAWordGroup(strings.Join(wordGroup, ""))
		//fmt.Println(*idWordGroup, string(bolB))
		idx.fileSystem.GetAttributeGroupWord().AddGroupWordsOfAttribute(idAttribute, idWordGroup)

		//println("DOCUMENTO GRUPO", idDocument, *idWordGroup)
		idDocument := document.GetID()
		idx.fileSystem.GetGroupWordDocument().AddGroupWordDocument(idWordGroup, &idDocument)
		//println(*idWordGroup, idDocument)

	}

	//	println("==================================================")
	//	println(idx.fileSystem.GetAttributeWord().ToJson())
	//	println("ATT MAP", idx.fileSystem.GetAttributeMap().ToJson())
	//	println("WORD MAP", idx.fileSystem.GetWordMap().ToJson())
	//	println("ATT WORD", idx.fileSystem.GetAttributeWord().ToJson())
	//	println("WORD GROUP MAP", idx.fileSystem.GetWordGroupMap().ToJson())
	//	println("DOCUMENT GROUP", idx.fileSystem.GetGroupWordDocument().ToJson())

	//println(idx.fileSystem.GetWordGroupMap().ToJson())
	//println(idx.fileSystem.GetGroupWordDocument().ToJson())

}
