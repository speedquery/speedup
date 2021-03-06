package filesystem

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
	"speedup/collection"
	"strconv"
	"sync"
	"time"
)

/*
FileSystem tem a função de fazer o gerenciamento de todos os indices
criados
*/

const (
	inverted        = "invertedindex"
	groupdocument   = "groupdocument"
	invertedworddoc = "invertedworddoc"
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

type FileSystem struct {
	wordmap            *WordMap
	attributeMap       *AttributeMap
	attributeWord      *AttributeWord
	wordGroupMap       *WordGroupMap
	attributeGroupWord *AttributeGroupWord
	groupWordDocument  *GroupWordDocument
	documentGroupWord  *DocumentGroupWord
	serialization      *Serialization
	wordDocument       *WordDocument
	Configuration      map[string]string
	FileSystem         *FileSystem
}

func (self *FileSystem) CreateFileSystem(nameFileSystem string, workFolder string) *FileSystem {

	self.Configuration = make(map[string]string)
	self.Configuration["nameFileSystem"] = nameFileSystem
	self.Configuration["path"] = workFolder
	self.Configuration["fileSystemFolder"] = workFolder + GetBar() + nameFileSystem

	self.FileSystem = self
	self.createWorkFolder()
	self.wordmap = new(WordMap).InitWordMap()
	self.attributeMap = new(AttributeMap).IniAttributeMap()
	self.attributeWord = new(AttributeWord).InitAttributeWord()
	self.wordGroupMap = new(WordGroupMap).IniWordGroupMap()
	self.attributeGroupWord = new(AttributeGroupWord).InitAttributeGroupWord()
	self.wordDocument = new(WordDocument).InitWordDocument(self.Configuration["fileSystemFolder"], self)
	self.groupWordDocument = new(GroupWordDocument).InitGroupWordDocument(self.Configuration["fileSystemFolder"])
	self.documentGroupWord = new(DocumentGroupWord).InitDocumentGroupWord(self.Configuration["fileSystemFolder"])

	self.serialization = new(Serialization).CreateSerialization(self)

	return self
}

func (self *FileSystem) createWorkFolder() {

	path := self.Configuration["fileSystemFolder"]

	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.Mkdir(path, 0777)

		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {

		path := path + GetBar() + inverted

		if _, err := os.Stat(path); os.IsNotExist(err) {

			os.Mkdir(path, 0777)

			println("CREATE INDEX:", path)
		}
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {

		path := path + GetBar() + groupdocument

		if _, err := os.Stat(path); os.IsNotExist(err) {

			os.Mkdir(path, 0777)

			println("CREATE INDEX:", path)
		}
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {

		path := path + GetBar() + invertedworddoc

		if _, err := os.Stat(path); os.IsNotExist(err) {

			os.Mkdir(path, 0777)

			println("CREATE INDEX:", path)
		}
	}

}

func (self *FileSystem) GetWordMap() *WordMap {
	return self.wordmap
}

func (self *FileSystem) GetAttributeMap() *AttributeMap {
	return self.attributeMap
}

func (self *FileSystem) GetAttributeWord() *AttributeWord {
	return self.attributeWord
}

func (self *FileSystem) GetWordGroupMap() *WordGroupMap {
	return self.wordGroupMap
}

func (self *FileSystem) GetAttributeGroupWord() *AttributeGroupWord {
	return self.attributeGroupWord
}

func (self *FileSystem) GetGroupWordDocument() *GroupWordDocument {
	return self.groupWordDocument
}

func (self *FileSystem) GetDocumentGroupWord() *DocumentGroupWord {
	return self.documentGroupWord
}

func (self *FileSystem) GetWordDocument() *WordDocument {
	return self.wordDocument
}

const (
	attributeMapFile      = "attmp.json"
	wordMapFile           = "wordmp.json"
	wordGroupMapFile      = "wordgpmp.json"
	attributeGroupWord    = "attgroupword-index.json"
	groupWordDocumentFile = "groupworddoc-index.json"
	documentGroupWordFile = "docgroupword-index.json"
	worddocumentsFile     = "worddocuments-index.json"
	attributeWordFile     = "attributeWord-index.json"
)

type Serialization struct {
	filesystem *FileSystem
}

func (self *Serialization) getBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

func (self *Serialization) CreateSerialization(filesystem *FileSystem) *Serialization {

	self.filesystem = filesystem

	var wg sync.WaitGroup

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeWordMap()

	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeAttributeMap()
	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeWordGroupMap()
	}(func() { wg.Done() })

	wg.Wait()

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeAttributeWord()

	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeWordDocuments()

	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeAttributeGroupWord()

	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeGroupWordDocument()

	}(func() { wg.Done() })

	wg.Add(1)
	go func(onClose func()) {
		defer onClose()
		self.DeSerealizeDocumentGroupWord()

	}(func() { wg.Done() })

	wg.Wait()

	//println(self.filesystem.GetAttributeMap().GetID())
	//println(self.filesystem.GetWordMap().GetID())
	//println(self.filesystem.GetWordGroupMap().GetID())

	println("SUCESS: DESEREALLIZATION:", self.filesystem.Configuration["nameFileSystem"])

	go func() {

		for {

			time.Sleep(time.Minute)

			go self.SerealizeAttributeMap()
			go self.SerealizeWordMap()
			go self.SerealizeWordGroupMap()
			go self.SerealizeWordDocuments()
			go self.SerealizeAttributeGroupWord()
			go self.SerealizeGroupWordDocument()
			go self.SerealizeDocumentGroupWord()
			go self.SerealizeAttributeWord()
		}

	}()

	return self

}

func (self *Serialization) createFile(nameFile string) *os.File {

	//*bufio.Writer

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + nameFile

	file, err := os.Create(path)

	//file.Close()

	if err != nil {
		panic(err)
	}

	return file
}

func (self *Serialization) GetLastID(fileName string) uint {
	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + "id-" + fileName

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("ERROR: NOT FOUND ID IN " + "id-" + fileName)
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]uint)
	json.Unmarshal([]byte(jsonString), &fields)

	println(fields["maxid"], fileName)

	return fields["maxid"]

}

func (self *Serialization) SerealizeAttributeMap() {

	json := self.filesystem.GetAttributeMap().ToJson()

	openedFile := self.createFile(attributeMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()

	json = self.filesystem.GetAttributeMap().ToJsonID()

	openedFile = self.createFile("id-" + attributeMapFile)
	bufferedWriter = bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()

}

func (self *Serialization) DeSerealizeAttributeMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + attributeMapFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	lastID := self.GetLastID(attributeMapFile)
	self.filesystem.GetAttributeMap().SetNewMap(lastID, newMap)

}

func (self *Serialization) SerealizeWordMap() {

	json := self.filesystem.GetWordMap().ToJson()
	openedFile := self.createFile(wordMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()

	json = self.filesystem.GetWordMap().ToJsonID()
	openedFile = self.createFile("id-" + wordMapFile)
	bufferedWriter = bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()

	//println("ID", json)

}

func (self *Serialization) DeSerealizeWordMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + wordMapFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	//scanner := bufio.NewScanner(openedFile)

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	//var jsonString string

	//for scanner.Scan() {
	//	println(scanner.Text())
	//	jsonString = scanner.Text()
	//}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	lastID := self.GetLastID(wordMapFile)

	self.filesystem.GetWordMap().SetNewMap(lastID, newMap)

}

func (self *Serialization) SerealizeWordGroupMap() {

	json := self.filesystem.GetWordGroupMap().ToJson()

	openedFile := self.createFile(wordGroupMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()

	json = self.filesystem.GetWordGroupMap().ToJsonID()

	openedFile = self.createFile("id-" + wordGroupMapFile)
	bufferedWriter = bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeWordGroupMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + wordGroupMapFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	lastID := self.GetLastID(wordGroupMapFile)
	self.filesystem.GetWordGroupMap().SetNewMap(lastID, newMap)

}

func (self *Serialization) SerealizeAttributeGroupWord() {

	json := self.filesystem.GetAttributeGroupWord().ToJson()

	openedFile := self.createFile(attributeGroupWord)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeAttributeGroupWord() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + attributeGroupWord

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[*uint]*collection.Set)

	for key, value := range fields {

		temp, err := strconv.Atoi(key)

		if err != nil {
			panic(err)
		}

		idAttribute := self.filesystem.GetAttributeMap().GetPointer(uint(temp))

		if idAttribute == nil {
			panic("PANIC: ID NÃO ENCONTRADO EM AttributeMap")
		}

		val := value.([]interface{})
		data := new(collection.Set).NewSet()

		for _, vl := range val {
			idWordGroup := self.filesystem.GetWordGroupMap().GetPointer(uint(vl.(float64)))

			if idWordGroup == nil {
				panic("PANIC: ID NÃO ENCONTRADO EM WordGroupMap")
			}

			data.Add(idWordGroup)
		}

		newMap[idAttribute] = data
	}

	self.filesystem.GetAttributeGroupWord().SetNewMap(newMap)

}

func (self *Serialization) SerealizeGroupWordDocument() {

	json := self.filesystem.GetGroupWordDocument().ToJson()

	openedFile := self.createFile(groupWordDocumentFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeGroupWordDocument() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + groupWordDocumentFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[*uint]*collection.Set)

	for key, _ := range fields {

		temp, err := strconv.Atoi(key)

		if err != nil {
			panic(err)
		}

		idGroup := self.filesystem.GetWordGroupMap().GetPointer(uint(temp))

		if idGroup == nil {
			panic("PANIC: ID NÃO ENCONTRADO EM WordGroupMap")
		}

		data := new(collection.Set).NewSet()

		newMap[idGroup] = data
	}

	self.filesystem.GetGroupWordDocument().SetNewMap(newMap)

}

func (self *Serialization) SerealizeDocumentGroupWord() {

	json := self.filesystem.GetDocumentGroupWord().ToJson()

	openedFile := self.createFile(documentGroupWordFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeDocumentGroupWord() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + documentGroupWordFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[*uint]*collection.Set)

	for key, _ := range fields {

		temp, err := strconv.Atoi(key)

		if err != nil {
			panic(err)
		}

		idDocument := uint(temp)
		data := new(collection.Set).NewSet()

		newMap[&idDocument] = data
	}

	self.filesystem.GetDocumentGroupWord().SetNewMap(newMap)

}

func (self *Serialization) SerealizeWordDocuments() {

	json := self.filesystem.GetWordDocument().ToJson()

	openedFile := self.createFile(worddocumentsFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeWordDocuments() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + worddocumentsFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[*uint]*collection.Set)

	for key, _ := range fields {

		temp, err := strconv.Atoi(key)

		if err != nil {
			panic(err)
		}

		idWord := self.filesystem.GetWordMap().GetPointKey(uint(temp))

		if idWord == nil {
			panic("NAO ENCONTRADO O ID EM WORDMAP")
		}

		data := new(collection.Set).NewSet()
		newMap[idWord] = data
	}

	self.filesystem.GetWordDocument().SetNewMap(newMap, self.filesystem)

}

func (self *Serialization) SerealizeAttributeWord() {

	json := self.filesystem.GetAttributeWord().ToJson()

	openedFile := self.createFile(attributeWordFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeAttributeWord() {

	path := self.filesystem.Configuration["fileSystemFolder"] + GetBar() + attributeWordFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	jsonString, err := ioutil.ReadAll(openedFile)

	if err != nil {
		panic(err)
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[*uint]*collection.Set)

	for key, value := range fields {

		tempkey, err := strconv.Atoi(key)

		if err != nil {
			panic(err)
		}

		idAttribute := self.filesystem.GetAttributeMap().GetPointer(uint(tempkey))

		if idAttribute == nil {
			panic("NAO ENCONTRADO O ID EM AttributeMap")
		}

		data := new(collection.Set).NewSet()
		arr := value.([]interface{})

		for _, value := range arr {

			tempkey := value.(float64) // := strconv.Atoi()

			if err != nil {
				panic(err)
			}

			idWord := self.filesystem.GetWordMap().GetPointKey(uint(tempkey))

			if idWord == nil {
				panic("ID NAO ENCONTRADO EM WordMap")
			}

			data.Add(idWord)
		}

		newMap[idAttribute] = data
	}

	self.filesystem.GetAttributeWord().SetNewMap(newMap)
	//println(self.filesystem.GetAttributeWord().ToJson())
}
