package filesystem

import (
	"bufio"
	"encoding/json"
	"os"
	"runtime"
	"speedup/collection"
	"strconv"
	"time"
)

/*
FileSystem tem a função de fazer o gerenciamento de todos os indices
criados
*/

const (
	inverted = "invertedindex"
)

type FileSystem struct {
	wordmap            *WordMap
	attributeMap       *AttributeMap
	attributeWord      *AttributeWord
	wordGroupMap       *WordGroupMap
	attributeGroupWord *AttributeGroupWord
	groupWordDocument  *GroupWordDocument
	serialization      *Serialization
	Configuration      map[string]string
}

func (self *FileSystem) getBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

func (self *FileSystem) CreateFileSystem(nameFileSystem string, workFolder string) *FileSystem {

	self.Configuration = make(map[string]string)
	self.Configuration["nameFileSystem"] = nameFileSystem
	self.Configuration["path"] = workFolder
	self.Configuration["fileSystemFolder"] = workFolder + self.getBar() + nameFileSystem

	self.createWorkFolder()
	self.wordmap = new(WordMap).InitWordMap()
	self.attributeMap = new(AttributeMap).IniAttributeMap()
	self.attributeWord = new(AttributeWord).InitAttributeWord()
	self.wordGroupMap = new(WordGroupMap).IniWordGroupMap()
	self.attributeGroupWord = new(AttributeGroupWord).InitAttributeGroupWord()
	self.groupWordDocument = new(GroupWordDocument).InitGroupWordDocument(self.Configuration["fileSystemFolder"])

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

		path = path + self.getBar() + inverted

		if _, err := os.Stat(path); os.IsNotExist(err) {

			os.Mkdir(path, 0777)

			//if err != nil {
			//	panic(err)
			//}

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

const (
	attributeMapFile      = "attmp.json"
	wordMapFile           = "wordmp.json"
	wordGroupMapFile      = "wordgpmp.json"
	attributeGroupWord    = "attgroupword-index.json"
	groupWordDocumentFile = "groupworddoc-index.json"
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

	self.DeSerealizeAttributeMap()
	self.DeSerealizeWordMap()
	self.DeSerealizeWordGroupMap()
	self.DeSerealizeAttributeGroupWord()
	self.DeSerealizeGroupWordDocument()

	println("SUCESS: DESEREALLIZATION:", self.filesystem.Configuration["nameFileSystem"])

	go func() {

		for {

			time.Sleep(time.Minute)

			go self.SerealizeAttributeMap()
			go self.SerealizeWordMap()
			go self.SerealizeWordGroupMap()
			go self.SerealizeAttributeGroupWord()
			go self.SerealizeGroupWordDocument()
		}

	}()

	return self

}

func (self *Serialization) createFile(nameFile string) *os.File {

	//*bufio.Writer

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + nameFile

	file, err := os.Create(path)

	//file.Close()

	if err != nil {
		panic(err)
	}

	return file
}

func (self *Serialization) SerealizeAttributeMap() {

	json := self.filesystem.GetAttributeMap().ToJson()

	openedFile := self.createFile(attributeMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeAttributeMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + attributeMapFile

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

	scanner := bufio.NewScanner(openedFile)

	var jsonString string

	for scanner.Scan() {
		jsonString = scanner.Text()
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	self.filesystem.GetAttributeMap().SetNewMap(newMap)

}

func (self *Serialization) SerealizeWordMap() {

	json := self.filesystem.GetWordMap().ToJson()

	openedFile := self.createFile(wordMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeWordMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + wordMapFile

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

	scanner := bufio.NewScanner(openedFile)

	var jsonString string

	for scanner.Scan() {
		jsonString = scanner.Text()
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	self.filesystem.GetWordMap().SetNewMap(newMap)

}

func (self *Serialization) SerealizeWordGroupMap() {

	json := self.filesystem.GetWordGroupMap().ToJson()

	openedFile := self.createFile(wordGroupMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *Serialization) DeSerealizeWordGroupMap() {

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + wordGroupMapFile

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

	scanner := bufio.NewScanner(openedFile)

	var jsonString string

	for scanner.Scan() {
		jsonString = scanner.Text()
	}

	openedFile.Close()

	fields := make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &fields)

	newMap := make(map[string]*uint)
	for key, value := range fields {
		data := uint(value.(float64))
		newMap[key] = &data
	}

	self.filesystem.GetWordGroupMap().SetNewMap(newMap)

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

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + attributeGroupWord

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

	scanner := bufio.NewScanner(openedFile)

	var jsonString string

	for scanner.Scan() {
		jsonString = scanner.Text()
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

	path := self.filesystem.Configuration["fileSystemFolder"] + self.getBar() + groupWordDocumentFile

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	openedFile, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(openedFile)

	var jsonString string

	for scanner.Scan() {
		jsonString = scanner.Text()
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
