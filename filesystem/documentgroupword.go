package filesystem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"speedup/collection"
	"sync"
	"time"
)

func (self *DocumentGroupWord) getBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

type DocumentGroupWord struct {
	//groupWordDocument map[*uint]*bufio.Writer
	//control           map[*uint]uint
	groupWordDocument map[*uint]*collection.Set
	someMapMutex      sync.RWMutex
	folder            string
	qtd               uint
}

func (self *DocumentGroupWord) SetNewMap(newMap map[*uint]*collection.Set) *DocumentGroupWord {
	self.groupWordDocument = newMap
	return self
}

func (self *DocumentGroupWord) InitDocumentGroupWord(fileSystemFolder string) *DocumentGroupWord {

	self.someMapMutex = sync.RWMutex{}
	self.groupWordDocument = make(map[*uint]*collection.Set)
	self.folder = fileSystemFolder + self.getBar() + groupdocument
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

			time.Sleep(time.Minute)

			data = self.Clone()

			for key, value := range data {

				path := self.folder + self.getBar() + fmt.Sprintf("%v", key) + ".txt"

				openedFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
				if err != nil {
					panic(err)
				}

				bufferedWriter := bufio.NewWriter(openedFile)

				for _, vl := range value {
					bufferedWriter.WriteString(fmt.Sprintf("%v", vl) + "\r\n")
				}

				bufferedWriter.Flush()

				//func(key uint, value []uint) {
				//
				//}(key, value)

				openedFile.Close()

			}

			println("ESCREVEU DOCUMENT GROUP")

			runtime.GC()

			//println("Limpou memoria", len(tm))

		}

	}()

	return self
}

func (self *DocumentGroupWord) createFile(name uint) {

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
func (gw *GroupWordDocument) GetIdGroupWord(idDocument *uint) *collection.Set {

	idGroups := gw.groupWordDocument[idDocument]
	return idGroups

}
**/

func (self *DocumentGroupWord) Clone() map[uint][]uint {

	temp := make(map[uint][]uint)

	self.someMapMutex.Lock()

	for key, value := range self.groupWordDocument {

		if value.Size() > 0 {

			data := make([]uint, 0)

			for vl, _ := range value.GetSet() {
				data = append(data, *vl)
			}

			temp[*key] = data
			self.groupWordDocument[key] = value.NewSet()
		}
	}

	self.someMapMutex.Unlock()

	//println("Clonou")

	return temp
}

func (self *DocumentGroupWord) Add(idGroup, idDocument *uint) (*collection.Set, bool) {

	data, exist := self.Get(idGroup)

	self.someMapMutex.Lock()
	if !exist || data == nil {

		self.createFile(*idGroup)

		data = new(collection.Set).NewSet()
		data.Add(idDocument)
		self.groupWordDocument[idGroup] = data
	} else {
		data.Add(idDocument)
	}
	self.someMapMutex.Unlock()
	return data, exist
}

func (self *DocumentGroupWord) Get(idGroup *uint) (*collection.Set, bool) {
	self.someMapMutex.Lock()
	data, exist := self.groupWordDocument[idGroup]
	self.someMapMutex.Unlock()

	return data, exist
}

func (self *DocumentGroupWord) AddDocumentGroupWord(idGroup, idDocument *uint) {

	self.Add(idGroup, idDocument)

}

func (self *DocumentGroupWord) ToJson() string {

	temp := make(map[uint]bool)

	self.someMapMutex.Lock()

	for key, _ := range self.groupWordDocument {

		temp[*key] = true
	}

	self.someMapMutex.Unlock()

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
