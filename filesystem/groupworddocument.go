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

func (self *GroupWordDocument) getBar() string {

	var bar string

	if runtime.GOOS == "windows" {
		bar = "\\"
	} else {
		bar = "/"
	}

	return bar

}

type GroupWordDocument struct {
	//groupWordDocument map[*uint]*bufio.Writer
	//control           map[*uint]uint
	groupWordDocument map[*uint]*collection.Set
	someMapMutex      sync.RWMutex
	folder            string
	qtd               uint
}

const (
	invertedFolder = "invertedindex"
)

func (self *GroupWordDocument) SetNewMap(newMap map[*uint]*collection.Set) *GroupWordDocument {
	self.groupWordDocument = newMap
	return self
}

func (self *GroupWordDocument) InitGroupWordDocument(fileSystemFolder string) *GroupWordDocument {

	self.someMapMutex = sync.RWMutex{}
	self.groupWordDocument = make(map[*uint]*collection.Set)
	self.folder = fileSystemFolder + self.getBar() + invertedFolder
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

			runtime.GC()

			//println("Limpou memoria", len(tm))

		}

	}()

	return self
}

func (self *GroupWordDocument) createFile(name uint) {

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

func (self *GroupWordDocument) Clone() map[uint][]uint {

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

func (self *GroupWordDocument) Add(idGroup, idDocument *uint) (*collection.Set, bool) {

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

func (self *GroupWordDocument) Get(idGroup *uint) (*collection.Set, bool) {
	self.someMapMutex.Lock()
	data, exist := self.groupWordDocument[idGroup]
	self.someMapMutex.Unlock()

	return data, exist
}

func (self *GroupWordDocument) AddGroupWordDocument(idGroup, idDocument *uint) {

	self.Add(idGroup, idDocument)

}

func (self *GroupWordDocument) ToJson() string {

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
