package filesystem

import (
	"bufio"
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
	**/
	if _, err := os.Stat(self.folder); os.IsNotExist(err) {
		os.Mkdir(self.folder, 0777)

		if err != nil {
			panic(err)
		}

		println("CREATE INDEX:", self.folder)

	}

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

func (gw *GroupWordDocument) Clone() map[uint][]uint {

	temp := make(map[uint][]uint)

	gw.someMapMutex.Lock()

	for key, value := range gw.groupWordDocument {

		if value.Size() > 0 {

			data := make([]uint, 0)

			for vl, _ := range value.GetSet() {
				data = append(data, *vl)
			}

			temp[*key] = data
			gw.groupWordDocument[key] = value.NewSet()
		}
	}

	gw.someMapMutex.Unlock()

	//println("Clonou")

	return temp
}

func (gw *GroupWordDocument) Add(idGroup, idDocument *uint) (*collection.Set, bool) {

	data, exist := gw.Get(idGroup)

	gw.someMapMutex.Lock()
	if !exist || data == nil {

		gw.createFile(*idGroup)

		data = new(collection.Set).NewSet()
		data.Add(idDocument)
		gw.groupWordDocument[idGroup] = data
	} else {
		data.Add(idDocument)
	}
	gw.someMapMutex.Unlock()
	return data, exist
}

func (gw *GroupWordDocument) Get(idGroup *uint) (*collection.Set, bool) {
	gw.someMapMutex.Lock()
	data, exist := gw.groupWordDocument[idGroup]
	gw.someMapMutex.Unlock()

	return data, exist
}

func (gw *GroupWordDocument) AddGroupWordDocument(idGroup, idDocument *uint) {

	gw.Add(idGroup, idDocument)

}

/**
func (gw *GroupWordDocument) ToJson() string {

	temp := make(map[uint][]*uint)

	for key, values := range gw.groupWordDocument {

		words := make([]*uint, 0)

		for key, _ := range values.GetSet() {
			words = append(words, key)
		}

		temp[*key] = words
	}

	data, err := json.Marshal(temp)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(data)

}
**/
