package indexwriter

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"speedup/collection"
	doc "speedup/document"
	fs "speedup/filesystem"
	"speedup/wordprocess/stringprocess"
	"strings"
	"sync"
	"time"
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
	fileSystem            *fs.FileSystem
	mapdeleteDocumentBulk *collection.SetUint
	someMapMutex          sync.RWMutex
}

func (self *IndexWriter) CreateIndex(fileSystem *fs.FileSystem) *IndexWriter {
	self.fileSystem = fileSystem
	self.mapdeleteDocumentBulk = new(collection.SetUint).NewSet()

	//go self.threadDeletDocumentBulk()

	return self
}

func (self *IndexWriter) IndexDocument(document *doc.Document, bulk bool) {

	tmp := document.GetID()
	var idDocument *uint = &tmp
	var wg sync.WaitGroup

	for attribute, value := range document.GetMap() {

		idAttribute := self.fileSystem.GetAttributeMap().AddAttribute(attribute)

		formatedValue := fmt.Sprintf("%v", value)
		words := strings.Split(formatedValue, " ")

		wordGroup := make([]string, 0) //list.New()

		for _, word := range words {
			//wg.Add(1)

			newWord := stringprocess.ProcessWord(word)
			idword := self.fileSystem.GetWordMap().AddWord(newWord)
			//idx.fileSystem.GetAttributeWord().AddWordsOfAttribute(idAttribute, idword)
			wordGroup = append(wordGroup, fmt.Sprint(*idword))
		}

		idWordGroup := self.fileSystem.GetWordGroupMap().AddAWordGroup(strings.Join(wordGroup, ""))

		wg.Add(1)
		go func(idAttribute *uint, idWordGroup *uint, onClose func()) {
			defer onClose()
			self.fileSystem.GetAttributeGroupWord().AddGroupWordsOfAttribute(idAttribute, idWordGroup)
		}(idAttribute, idWordGroup, func() { wg.Done() })

		wg.Add(1)
		go func(idWordGroup, idDocument *uint, bulk bool, onClose func()) {
			defer onClose()
			self.fileSystem.GetGroupWordDocument().AddGroupWordDocument(idWordGroup, idDocument, bulk)
		}(idWordGroup, idDocument, bulk, func() { wg.Done() })

		wg.Add(1)
		go func(idDocument, idWordGroup *uint, bulk bool, onClose func()) {
			defer onClose()
			self.fileSystem.GetDocumentGroupWord().AddDocumentGroupWord(idDocument, idWordGroup, bulk)
		}(idDocument, idWordGroup, bulk, func() { wg.Done() })

	}

	wg.Wait()

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

func (self *IndexWriter) UpdateDocument(document *doc.Document) bool {

	sucess := self.DeleteDocument(document.GetID())

	if sucess {
		self.IndexDocument(document, false)
	}

	return sucess
}

func (self *IndexWriter) DeleteDocumentBulk(idDocument uint) {

	self.someMapMutex.Lock()
	self.mapdeleteDocumentBulk.Add(idDocument)
	self.someMapMutex.Unlock()

}

func (self *IndexWriter) threadDeletDocumentBulk() {

	println("DELETE BULK")

	for {
		time.Sleep(time.Minute)

		documents := self.mapdeleteDocumentBulk.Clone()

		if len(documents) > 0 {

			for document, _ := range documents {
				self.DeleteDocument(document)
			}

			println("DELETOU UM LOTE", len(documents))

		} else {
			println("=========== CONCLUIU INDEX BULK ===========")
		}
	}
}

func (self *IndexWriter) DeleteDocument(idDocument uint) bool {

	path := self.fileSystem.Configuration["fileSystemFolder"] + GetBar() + groupdocument + GetBar() + fmt.Sprintf("%v", idDocument) + ".txt"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

<<<<<<< HEAD
	file, err := os.Open(path)
	if err != nil {
		panic(err)
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

			go func(idGroup string, idDocument uint, onClose func()) {

				defer onClose()

				path := self.fileSystem.Configuration["fileSystemFolder"] + GetBar() + inverted + GetBar() + fmt.Sprintf("%v", idGroup) + ".txt"

				if _, err := os.Stat(path); os.IsNotExist(err) {
					return
				}

				file, err := os.Open(path)
				if err != nil {
					panic(err)
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
					// else {
					//	println("ENCONTROU:", idDocument, idGroup)
					//}

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

			}(idGroup, idDocument, func() { wg.Done() })

		}

		wg.Wait()

	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {

		err = os.Remove(path)

		if err != nil {
			panic(err)
		}

	}
=======
	println(idx.fileSystem.GetWordMap().ToJson())
	println(idx.fileSystem.GetAttributeWord().ToJson())
>>>>>>> 4904f70da0c7e5c04f9c4351f39631fe978b1cca

	return true
}
