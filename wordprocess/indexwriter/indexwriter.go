package indexwriter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"speedup/collection"
	doc "speedup/document"
	fs "speedup/filesystem"
	"speedup/wordprocess/stringprocess"
	"strings"
	"sync"
)

func GetBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

const (
	inverted      = "invertedindex"
	groupdocument = "groupdocument"
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

	tmp := document.GetID()
	var idDocument *uint = &tmp

	for attribute, value := range document.GetMap() {

		idAttribute := idx.fileSystem.GetAttributeMap().AddAttribute(attribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := make([]string, 0) //list.New()

		for _, word := range words {

			newWord := stringprocess.ProcessWord(word)
			idword := idx.fileSystem.GetWordMap().AddWord(newWord)
			//idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)
			wordGroup = append(wordGroup, fmt.Sprint(*idword))
		}

		idWordGroup := idx.fileSystem.GetWordGroupMap().AddAWordGroup(strings.Join(wordGroup, ""))
		idx.fileSystem.GetAttributeGroupWord().AddGroupWordsOfAttribute(idAttribute, idWordGroup)
		idx.fileSystem.GetGroupWordDocument().AddGroupWordDocument(idWordGroup, idDocument)
		idx.fileSystem.GetDocumentGroupWord().AddDocumentGroupWord(idDocument, idWordGroup)
	}

	//document = document.DeleteMemoryDocument()
	//document = nil

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

func (self *IndexWriter) DeleteDocument(idDocument uint) bool {

	path := self.fileSystem.Configuration["fileSystemFolder"] + GetBar() + groupdocument + GetBar() + fmt.Sprintf("%v", idDocument) + ".txt"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	//println(path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	groupsDocuments := new(collection.StrSet).NewSet()

	for scanner.Scan() {
		groupsDocuments.Add(scanner.Text())
	}

	file.Close()

	if groupsDocuments.Size() > 0 {

		var wg sync.WaitGroup

		for idGroup, _ := range groupsDocuments.GetSet() {

			wg.Add(1)

			go func(idGroup string, onClose func()) {

				defer onClose()

				path := self.fileSystem.Configuration["fileSystemFolder"] + GetBar() + inverted + GetBar() + fmt.Sprintf("%v", idGroup) + ".txt"

				file, err := os.Open(path)
				if err != nil {
					log.Fatal(err)
				}

				set := new(collection.StrSet).NewSet()
				scanner := bufio.NewScanner(file)

				//println("Grupo", idGroup)
				for scanner.Scan() {
					//println("Documento", scanner.Text())
					value := scanner.Text()
					if value != fmt.Sprintf("%v", idDocument) {
						set.Add(scanner.Text())
					}

				}

				if _, err := os.Stat(path); !os.IsNotExist(err) {
					os.Remove(path)

					if err != nil {
						panic(err)
					}
				}

				//cria arquivo
				_, err = os.Create(path)

				if err != nil {
					panic(err)
				}

				//abre arquivo
				openedFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
				if err != nil {
					panic(err)
				}

				bufferedWriter := bufio.NewWriter(openedFile)

				for vl, _ := range set.GetSet() {
					bufferedWriter.WriteString(vl + "\r\n")
				}

				bufferedWriter.Flush()
				openedFile.Close()
			}(idGroup, func() { wg.Done() })

		}

		wg.Wait()

	}

	err = os.Remove(path)

	if err != nil {
		panic(err)
	}

	return true
}
