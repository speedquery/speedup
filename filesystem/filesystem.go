package filesystem

import (
	"bufio"
	"encoding/json"
	"os"
	"speedup/collection"
	"strconv"
)

/*
FileSystem tem a função de fazer o gerenciamento de todos os indices
criados
*/

const (
	attributeMapFile   = "attmp.json"
	wordMapFile        = "wordmp.json"
	wordGroupMapFile   = "wordgpmp.json"
	attributeGroupWord = "attgroupword-index.json"
)

type FileSystem struct {
	wordmap            *WordMap
	attributeMap       *AttributeMap
	attributeWord      *AttributeWord
	wordGroupMap       *WordGroupMap
	attributeGroupWord *AttributeGroupWord
	groupWordDocument  *GroupWordDocument
	nameFileSystem     string
	path               string
	configuration      map[string]string
}

func (self *FileSystem) CreateFileSystem(nameFileSystem string, workFolder string) *FileSystem {

	self.configuration = make(map[string]string)
	self.configuration["nameFileSystem"] = nameFileSystem
	self.configuration["path"] = workFolder
	self.configuration["fileSystemFolder"] = workFolder + "\\" + nameFileSystem

	self.wordmap = new(WordMap).InitWordMap()
	self.attributeMap = new(AttributeMap).IniAttributeMap()
	self.attributeWord = new(AttributeWord).InitAttributeWord()
	self.wordGroupMap = new(WordGroupMap).IniWordGroupMap()
	self.attributeGroupWord = new(AttributeGroupWord).InitAttributeGroupWord()
	self.groupWordDocument = new(GroupWordDocument).InitGroupWordDocument()

	self.createWorkFolder()

	self.DeSerealizeAttributeMap()
	self.DeSerealizeWordMap()
	self.DeSerealizeWordGroupMap()
	self.DeSerealizeAttributeGroupWord()

	return self
}

func (self *FileSystem) createWorkFolder() {

	if _, err := os.Stat(self.configuration["fileSystemFolder"]); os.IsNotExist(err) {
		os.Mkdir(self.configuration["fileSystemFolder"], 0777)
	}
}

func (self *FileSystem) createFile(nameFile string) *os.File {

	//*bufio.Writer

	path := self.configuration["fileSystemFolder"] + "\\" + nameFile

	file, err := os.Create(path)

	//file.Close()

	if err != nil {
		panic(err)
	}

	return file
}

func (self *FileSystem) SerealizeAttributeMap() {

	json := self.GetAttributeMap().ToJson()

	openedFile := self.createFile(attributeMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *FileSystem) DeSerealizeAttributeMap() {

	path := self.configuration["fileSystemFolder"] + "\\" + attributeMapFile

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

	self.GetAttributeMap().SetNewMap(newMap)

}

func (self *FileSystem) SerealizeWordMap() {

	json := self.GetWordMap().ToJson()

	openedFile := self.createFile(wordMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *FileSystem) DeSerealizeWordMap() {

	path := self.configuration["fileSystemFolder"] + "\\" + wordMapFile
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

	self.GetWordMap().SetNewMap(newMap)

}

func (self *FileSystem) SerealizeWordGroupMap() {

	json := self.GetWordGroupMap().ToJson()

	openedFile := self.createFile(wordGroupMapFile)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *FileSystem) DeSerealizeWordGroupMap() {

	path := self.configuration["fileSystemFolder"] + "\\" + wordGroupMapFile
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

	self.GetWordGroupMap().SetNewMap(newMap)

}

func (self *FileSystem) SerealizeAttributeGroupWord() {

	json := self.GetAttributeGroupWord().ToJson()

	openedFile := self.createFile(attributeGroupWord)
	bufferedWriter := bufio.NewWriter(openedFile)
	bufferedWriter.WriteString(json)
	bufferedWriter.Flush()
	openedFile.Close()
}

func (self *FileSystem) DeSerealizeAttributeGroupWord() {

	path := self.configuration["fileSystemFolder"] + "\\" + attributeGroupWord
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

		idAttribute := self.GetAttributeMap().GetPointer(uint(temp))

		if idAttribute == nil {
			panic("PANIC: ID NÃO ENCONTRADO EM AttributeMap")
		}

		val := value.([]interface{})
		data := new(collection.Set).NewSet()

		for _, vl := range val {
			idWordGroup := self.GetWordGroupMap().GetPointer(uint(vl.(float64)))

			if idWordGroup == nil {
				panic("PANIC: ID NÃO ENCONTRADO EM WordGroupMap")
			}

			data.Add(idWordGroup)
		}

		newMap[idAttribute] = data
	}

	self.GetAttributeGroupWord().SetNewMap(newMap)
	println("Serializou")

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
