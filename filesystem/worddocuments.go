package filesystem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"speedup/collection"
	"speedup/utils"
	"strings"
	"sync"
	"time"
)

func (self *WordDocument) getBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

type WordDocument struct {
	//WordDocument map[*uint]*bufio.Writer
	//control           map[*uint]uint
	WordDocument map[*uint]*collection.Set
	numberList   *collection.Set
	someMapMutex sync.RWMutex
	folder       string
	qtd          uint
	fileSystem   *FileSystem
}

func (self *WordDocument) SetNewMap(newMap map[*uint]*collection.Set, fileSystem *FileSystem) *WordDocument {
	self.WordDocument = newMap

	self.numberList = new(collection.Set).NewSet()

	self.fileSystem = fileSystem

	for k, _ := range self.WordDocument {

		word := self.fileSystem.GetWordMap().GetValue(k)

		if len(*word) == 0 || !utils.IsNumber(*word) {
			continue
		}

		self.numberList.Add(k)
	}

	return self
}

func (self *WordDocument) GetList() *collection.Set {

	return self.numberList

}

func (self *WordDocument) GetNumberList() []*uint {

	temp := make([]*uint, 0)

	for k, _ := range self.numberList.GetSet() {
		temp = append(temp, k)
	}

	return temp
}

func (self *WordDocument) InitWordDocument(fileSystemFolder string, fileSystem *FileSystem) *WordDocument {

	self.someMapMutex = sync.RWMutex{}
	self.WordDocument = make(map[*uint]*collection.Set)
	self.numberList = new(collection.Set).NewSet()

	self.folder = fileSystemFolder + self.getBar() + invertedworddoc
	self.qtd = 0
	/**
			if runtime.GOOS == "windows" {
				gw.folder = "C:/data2"
			} else {
				gw.folder = "/users/thiagorodrigues/documents/goteste"
			}

	if _, err := os.Stat(self.folder); os.IsNotExist(err) {

		println(self.folder)

		os.Mkdir(self.folder, 0777)

		if err != nil {
			panic(err)
		}

		println("CREATE INDEX:", self.folder)

	}
	**/

	go func() {

		var data map[uint][]uint

		for {

			time.Sleep(time.Second * 60)

			data = self.Clone()

			if len(data) > 0 {
				go self.WriterInFile(data)
			}
		}

	}()

	return self
}

func (self *WordDocument) WriterInFile(data map[uint][]uint) {

	var wg sync.WaitGroup

	for key, value := range data {

		path := self.folder + self.getBar() + fmt.Sprintf("%v", key) + ".txt"

		openedFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		bufferedWriter := bufio.NewWriter(openedFile)
		//wg.Add(1)
		for _, vl := range value {
			bufferedWriter.WriteString(strings.TrimSpace(fmt.Sprintf("%v", vl)) + "\r\n")
			/**
			go func(vl string, onClose func()) {
				defer onClose()
				bufferedWriter.WriteString(vl + "\r\n")
			}(fmt.Sprintf("%v", vl), func() { wg.Done() })
			**/
		}

		wg.Wait()

		bufferedWriter.Flush()

		openedFile.Close()

	}

}

func (self *WordDocument) createFile(name uint) {

	//*bufio.Writer

	path := self.folder + self.getBar() + fmt.Sprintf("%v", name) + ".txt"

	file, err := os.Create(path)

	file.Close()

	if err != nil {
		panic(err)
	}

	/**

	openedFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	//defer openedFile.Close()

	bufferedWriter := bufio.NewWriter(openedFile)

	return bufferedWriter
	*/ //

}

/**
func (gw *WordDocument) GetidWordWord(idDocument *uint) *collection.Set {

	idWords := gw.WordDocument[idDocument]
	return idWords

}
**/

func (self *WordDocument) Clone() map[uint][]uint {

	temp := make(map[uint][]uint)

	self.someMapMutex.Lock()

	for key, value := range self.WordDocument {

		if value.Size() > 0 {

			data := make([]uint, 0)

			for vl, _ := range value.GetSet() {
				data = append(data, *vl)
			}

			temp[*key] = data
			self.WordDocument[key] = value.NewSet()
		}
	}

	self.someMapMutex.Unlock()

	//println("Clonou")

	return temp
}

func (self *WordDocument) Add(idWord, idDocument *uint) (*collection.Set, bool) {

	data, exist := self.Get(idWord)

	self.someMapMutex.Lock()
	if !exist || data == nil {

		self.createFile(*idWord)

		data = new(collection.Set).NewSet()
		data.Add(idDocument)
		self.WordDocument[idWord] = data
	} else {
		data.Add(idDocument)
	}
	self.someMapMutex.Unlock()
	return data, exist
}

func (self *WordDocument) Get(idWord *uint) (*collection.Set, bool) {
	self.someMapMutex.Lock()
	data, exist := self.WordDocument[idWord]
	self.someMapMutex.Unlock()

	return data, exist
}

func (self *WordDocument) AddWordDocument(idWord, idDocument *uint, bulk bool) {

	self.Add(idWord, idDocument)

	if !bulk {
		dados := self.Clone()
		self.WriterInFile(dados)
	}

}

func (self *WordDocument) ToJson() string {

	temp := make(map[uint]bool)

	self.someMapMutex.Lock()

	for key, _ := range self.WordDocument {

		temp[*key] = true
	}

	self.someMapMutex.Unlock()

	data, err := json.Marshal(temp)

	if err != nil {
		panic(err)
	}

	return string(data)

}
