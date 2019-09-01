package serialization

import (
	"bufio"
	"encoding/json"
	"os"
	"runtime"
	"speedup/collection"
	"speedup/filesystem"
	"strconv"
	"time"
)

const (
	attributeMapFile   = "attmp.json"
	wordMapFile        = "wordmp.json"
	wordGroupMapFile   = "wordgpmp.json"
	attributeGroupWord = "attgroupword-index.json"
)

type Serialization struct {
	filesystem *filesystem.FileSystem
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

func (self *Serialization) CreateSerialization(filesystem *filesystem.FileSystem) *Serialization {

	self.filesystem = filesystem

	self.createWorkFolder()
	self.DeSerealizeAttributeMap()
	self.DeSerealizeWordMap()
	self.DeSerealizeWordGroupMap()
	self.DeSerealizeAttributeGroupWord()

	println("SUCESS: DESEREALLIZATION")

	go func() {

		for {

			time.Sleep(time.Minute)

			go self.SerealizeAttributeMap()
			go self.SerealizeWordMap()
			go self.SerealizeWordGroupMap()
			go self.SerealizeAttributeGroupWord()
		}

	}()

	return self

}

func (self *Serialization) createWorkFolder() {

	if _, err := os.Stat(self.filesystem.Configuration["fileSystemFolder"]); os.IsNotExist(err) {
		os.Mkdir(self.filesystem.Configuration["fileSystemFolder"], 0777)
	}
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
